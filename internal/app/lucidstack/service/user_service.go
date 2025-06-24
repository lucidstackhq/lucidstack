package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/field"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/request"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/response"
	"github.com/lucidstackhq/lucidstack/internal/pkg/auth"
	"github.com/lucidstackhq/lucidstack/internal/pkg/secret"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	organizationService *OrganizationService
	authenticator       *auth.Authenticator
}

func NewUserService(organizationService *OrganizationService, authenticator *auth.Authenticator) *UserService {
	return &UserService{organizationService: organizationService, authenticator: authenticator}
}

func (s *UserService) SignUp(ctx context.Context, request *request.UserSignUpRequest) (*model.User, error) {
	organizationExists, err := s.organizationService.NameExists(ctx, request.OrganizationName)

	if err != nil {
		return nil, err
	}

	if organizationExists {
		return nil, fmt.Errorf("organization %s already exists", request.OrganizationName)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:       request.Username,
		Password:       string(hashedPassword),
		Admin:          true,
		CreatorID:      "",
		OrganizationID: "",
	}

	err = mgm.Coll(user).CreateWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	organization, err := s.organizationService.Save(ctx, request.OrganizationName, request.BillingEmail, user.ID.Hex())

	if err != nil {
		return nil, err
	}

	user.OrganizationID = organization.ID.Hex()

	err = mgm.Coll(user).UpdateWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetToken(ctx context.Context, request *request.UserTokenRequest) (*response.UserTokenResponse, error) {

	organization, err := s.organizationService.GetByName(ctx, request.OrganizationName)

	if err != nil {
		return nil, err
	}

	user := &model.User{}

	err = mgm.Coll(user).FirstWithCtx(ctx, bson.M{
		"username":        request.Username,
		"organization_id": organization.ID.Hex(),
	}, user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("invalid username and password combination")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return nil, fmt.Errorf("invalid username and password")
	}

	token, err := s.authenticator.GenerateUserToken(user.ID.Hex(), user.OrganizationID, user.Admin)

	if err != nil {
		return nil, err
	}

	return &response.UserTokenResponse{Token: token}, nil
}

func (s *UserService) Get(ctx context.Context, userID string, organizationID string) (*model.User, error) {
	id, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	user := &model.User{}

	err = mgm.Coll(user).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	}, user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ChangePassword(ctx context.Context, userID string, request *request.UserPasswordChangeRequest, organizationID string) (*model.User, error) {
	user, err := s.Get(ctx, userID, organizationID)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	err = mgm.Coll(user).UpdateWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Add(ctx context.Context, request *request.UserAdditionRequest, creatorID string, organizationID string) (*response.UserPasswordResponse, error) {
	usernameExists, err := s.usernameExists(ctx, request.Username, organizationID)

	if err != nil {
		return nil, err
	}

	if usernameExists {
		return nil, fmt.Errorf("username %s is already taken", request.Username)
	}

	newPassword, err := secret.GenerateSecret(16)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:       request.Username,
		Password:       string(hashedPassword),
		Admin:          request.Admin,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(user).CreateWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	return &response.UserPasswordResponse{
		Password: newPassword,
		User:     user,
	}, nil
}

func (s *UserService) Delete(ctx context.Context, userID string, organizationID string) (*model.User, error) {
	user, err := s.Get(ctx, userID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(user).DeleteWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(ctx context.Context, organizationID string, page int64, size int64) ([]*model.User, error) {
	users := make([]*model.User, 0)

	err := mgm.Coll(&model.User{}).SimpleFindWithCtx(ctx, &users, bson.M{
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size*size))

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateAdmin(ctx context.Context, userID string, request *request.UserAdminUpdateRequest, organizationID string) (*model.User, error) {
	user, err := s.Get(ctx, userID, organizationID)

	if err != nil {
		return nil, err
	}

	user.Admin = request.Admin

	err = mgm.Coll(user).UpdateWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ResetPassword(ctx context.Context, userID string, organizationID string) (*response.UserPasswordResponse, error) {
	user, err := s.Get(ctx, userID, organizationID)

	if err != nil {
		return nil, err
	}

	newPassword, err := secret.GenerateSecret(16)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	err = mgm.Coll(user).UpdateWithCtx(ctx, user)

	if err != nil {
		return nil, err
	}

	return &response.UserPasswordResponse{
		Password: newPassword,
		User:     user,
	}, nil
}

func (s *UserService) usernameExists(ctx context.Context, username string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.User{}).CountDocuments(ctx, bson.M{"username": username, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
