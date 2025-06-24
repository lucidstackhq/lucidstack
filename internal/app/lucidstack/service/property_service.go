package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/field"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PropertyService struct {
	modelService *ModelService
}

func NewPropertyService(modelService *ModelService) *PropertyService {
	return &PropertyService{modelService: modelService}
}

func (s *PropertyService) Create(ctx context.Context, modelID string, request *request.PropertyCreationRequest, creatorID string, organizationID string) (*model.Property, error) {
	modelExists, err := s.modelService.Exists(ctx, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if !modelExists {
		return nil, fmt.Errorf("model not found")
	}

	nameExists, err := s.nameExists(ctx, request.Name, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("property %s already exists", request.Name)
	}

	property := &model.Property{
		ModelID:        modelID,
		Name:           request.Name,
		Description:    request.Description,
		DataSchema:     request.DataSchema,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(property).CreateWithCtx(ctx, property)

	if err != nil {
		return nil, err
	}

	return property, nil
}

func (s *PropertyService) List(ctx context.Context, modelID string, organizationID string, page int64, size int64) ([]*model.Property, error) {
	properties := make([]*model.Property, 0)

	err := mgm.Coll(&model.Property{}).SimpleFindWithCtx(ctx, &properties, bson.M{
		"model_id":        modelID,
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (s *PropertyService) Get(ctx context.Context, propertyID string, modelID string, organizationID string) (*model.Property, error) {
	id, err := primitive.ObjectIDFromHex(propertyID)

	if err != nil {
		return nil, err
	}

	property := &model.Property{}

	err = mgm.Coll(property).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"model_id":        modelID,
		"organization_id": organizationID,
	}, property)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("property not found")
	}

	if err != nil {
		return nil, err
	}

	return property, nil
}

func (s *PropertyService) Update(ctx context.Context, propertyID string, request *request.PropertyUpdateRequest, modelID string, organizationID string) (*model.Property, error) {
	property, err := s.Get(ctx, propertyID, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	property.Description = request.Description

	err = mgm.Coll(property).UpdateWithCtx(ctx, property)

	if err != nil {
		return nil, err
	}

	return property, nil
}

func (s *PropertyService) Delete(ctx context.Context, propertyID string, modelID string, organizationID string) (*model.Property, error) {
	property, err := s.Get(ctx, propertyID, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(property).DeleteWithCtx(ctx, property)

	if err != nil {
		return nil, err
	}

	return property, nil
}

func (s *PropertyService) nameExists(ctx context.Context, name string, modelID string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.Model{}).CountDocuments(ctx, bson.M{"name": name, "model_id": modelID, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
