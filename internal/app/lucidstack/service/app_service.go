package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/field"
	"github.com/kamva/mgm/v3/operator"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/request"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/response"
	"github.com/lucidstackhq/lucidstack/internal/pkg/secret"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppService struct {
}

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) Create(ctx context.Context, request *request.AppCreationRequest, creatorID string, organizationID string) (*model.App, error) {
	nameExists, err := s.nameExists(ctx, request.Name, organizationID)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("app %s already exists", request.Name)
	}

	appSecret, err := secret.GenerateSecret(128)

	if err != nil {
		return nil, err
	}

	app := &model.App{
		Name:           request.Name,
		Description:    request.Description,
		Secret:         appSecret,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(app).CreateWithCtx(ctx, app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppService) List(ctx context.Context, organizationID string, page int64, size int64) ([]*model.App, error) {
	apps := make([]*model.App, 0)

	err := mgm.Coll(&model.App{}).SimpleFindWithCtx(ctx, &apps, bson.M{
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (s *AppService) Get(ctx context.Context, appID string, organizationID string) (*model.App, error) {
	id, err := primitive.ObjectIDFromHex(appID)

	if err != nil {
		return nil, err
	}

	app := &model.App{}

	err = mgm.Coll(app).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	}, app)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("app not found")
	}

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppService) Update(ctx context.Context, appID string, request *request.AppUpdateRequest, organizationID string) (*model.App, error) {
	app, err := s.Get(ctx, appID, organizationID)

	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		count, err := mgm.Coll(app).CountDocuments(ctx, bson.M{
			field.ID: bson.M{
				operator.Ne: app.ID,
			},
			"name":            request.Name,
			"organization_id": organizationID,
		})

		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf("app %s already exists", request.Name)
		}

		app.Name = request.Name
	}

	app.Description = request.Description

	err = mgm.Coll(app).UpdateWithCtx(ctx, app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppService) Delete(ctx context.Context, appID string, organizationID string) (*model.App, error) {
	app, err := s.Get(ctx, appID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(app).DeleteWithCtx(ctx, app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppService) GetSecret(ctx context.Context, appID string, organizationID string) (*response.AppSecretResponse, error) {
	app, err := s.Get(ctx, appID, organizationID)

	if err != nil {
		return nil, err
	}

	return &response.AppSecretResponse{Secret: app.Secret}, nil
}

func (s *AppService) ResetSecret(ctx context.Context, appID string, organizationID string) (*response.AppSecretResponse, error) {
	app, err := s.Get(ctx, appID, organizationID)

	if err != nil {
		return nil, err
	}

	appSecret, err := secret.GenerateSecret(128)

	if err != nil {
		return nil, err
	}

	app.Secret = appSecret

	err = mgm.Coll(app).UpdateWithCtx(ctx, app)

	if err != nil {
		return nil, err
	}

	return &response.AppSecretResponse{Secret: appSecret}, nil
}

func (s *AppService) nameExists(ctx context.Context, name string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.App{}).CountDocuments(ctx, bson.M{"name": name, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
