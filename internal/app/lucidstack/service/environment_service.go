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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnvironmentService struct {
}

func NewEnvironmentService() *EnvironmentService {
	return &EnvironmentService{}
}

func (s *EnvironmentService) Create(ctx context.Context, request *request.EnvironmentCreationRequest, creatorID string, organizationID string) (*model.Environment, error) {

	nameExists, err := s.nameExists(ctx, request.Name, organizationID)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("environment %s already exists", request.Name)
	}

	environment := &model.Environment{
		Name:           request.Name,
		Description:    request.Description,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(environment).CreateWithCtx(ctx, environment)

	if err != nil {
		return nil, err
	}

	return environment, nil
}

func (s *EnvironmentService) List(ctx context.Context, organizationID string, page int64, size int64) ([]*model.Environment, error) {
	environments := make([]*model.Environment, 0)

	err := mgm.Coll(&model.Environment{}).SimpleFindWithCtx(ctx, &environments, bson.M{"organization_id": organizationID}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return environments, nil
}

func (s *EnvironmentService) Get(ctx context.Context, environmentID string, organizationID string) (*model.Environment, error) {
	id, err := primitive.ObjectIDFromHex(environmentID)

	if err != nil {
		return nil, err
	}

	environment := &model.Environment{}

	err = mgm.Coll(environment).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	}, environment)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("environment not found")
	}

	if err != nil {
		return nil, err
	}

	return environment, nil
}

func (s *EnvironmentService) Update(ctx context.Context, environmentID string, request *request.EnvironmentUpdateRequest, organizationID string) (*model.Environment, error) {
	environment, err := s.Get(ctx, environmentID, organizationID)

	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		count, err := mgm.Coll(environment).CountDocuments(ctx, bson.M{
			field.ID:          bson.M{operator.Ne: environment.ID},
			"name":            request.Name,
			"organization_id": organizationID,
		})

		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf("environment %s already exists", request.Name)
		}

		environment.Name = request.Name
	}

	environment.Description = request.Description

	err = mgm.Coll(environment).UpdateWithCtx(ctx, environment)

	if err != nil {
		return nil, err
	}

	return environment, nil
}

func (s *EnvironmentService) Delete(ctx context.Context, environmentID string, organizationID string) (*model.Environment, error) {
	environment, err := s.Get(ctx, environmentID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(environment).DeleteWithCtx(ctx, environment)

	if err != nil {
		return nil, err
	}

	return environment, nil
}

func (s *EnvironmentService) Exists(ctx context.Context, environmentID string, organizationID string) (bool, error) {
	id, err := primitive.ObjectIDFromHex(environmentID)

	if err != nil {
		return false, err
	}

	count, err := mgm.Coll(&model.Environment{}).CountDocuments(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *EnvironmentService) Search(ctx context.Context, query string, organizationID string, page int64, size int64) ([]*model.Environment, error) {
	if query == "" {
		return s.List(ctx, organizationID, page, size)
	}

	environments := make([]*model.Environment, 0)

	err := mgm.Coll(&model.Environment{}).SimpleFindWithCtx(ctx, &environments, bson.M{
		operator.Text: bson.M{
			"$search": query,
		},
		"organization_id": organizationID,
	})

	if err != nil {
		return nil, err
	}

	return environments, nil
}

func (s *EnvironmentService) nameExists(ctx context.Context, name string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.Environment{}).CountDocuments(ctx, bson.M{"name": name, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
