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

type ModelService struct {
}

func NewModelService() *ModelService {
	return &ModelService{}
}

func (s *ModelService) Create(ctx context.Context, request *request.ModelCreationRequest, creatorID string, organizationID string) (*model.Model, error) {
	nameExists, err := s.nameExists(ctx, request.Name, organizationID)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("model %s already exists", request.Name)
	}

	modelData := &model.Model{
		Name:           request.Name,
		Description:    request.Description,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(modelData).CreateWithCtx(ctx, modelData)

	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (s *ModelService) List(ctx context.Context, organizationID string, page int64, size int64) ([]*model.Model, error) {
	models := make([]*model.Model, 0)

	err := mgm.Coll(&model.Model{}).SimpleFindWithCtx(ctx, &models, bson.M{
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s *ModelService) Get(ctx context.Context, modelID string, organizationID string) (*model.Model, error) {
	id, err := primitive.ObjectIDFromHex(modelID)

	if err != nil {
		return nil, err
	}

	modelData := &model.Model{}

	err = mgm.Coll(modelData).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	}, modelData)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("model not found")
	}

	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (s *ModelService) Update(ctx context.Context, modelID string, request *request.ModelUpdateRequest, organizationID string) (*model.Model, error) {
	modelData, err := s.Get(ctx, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		count, err := mgm.Coll(modelData).CountDocuments(ctx, bson.M{
			field.ID: bson.M{
				operator.Ne: modelData.ID,
			},
			"name":            request.Name,
			"organization_id": organizationID,
		})

		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf("model %s already exists", request.Name)
		}

		modelData.Name = request.Name
	}

	modelData.Description = request.Description

	err = mgm.Coll(modelData).UpdateWithCtx(ctx, modelData)

	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (s *ModelService) Delete(ctx context.Context, modelID string, organizationID string) (*model.Model, error) {
	modelData, err := s.Get(ctx, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(modelData).DeleteWithCtx(ctx, modelData)

	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (s *ModelService) nameExists(ctx context.Context, name string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.Model{}).CountDocuments(ctx, bson.M{"name": name, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
