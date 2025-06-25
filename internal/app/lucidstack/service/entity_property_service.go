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

type EntityPropertyService struct {
	entityService   *EntityService
	propertyService *PropertyService
}

func NewEntityPropertyService(entityService *EntityService, propertyService *PropertyService) *EntityPropertyService {
	return &EntityPropertyService{entityService: entityService, propertyService: propertyService}
}

func (s *EntityPropertyService) Create(ctx context.Context, entityID string, propertyID string, request *request.EntityPropertyRequest, creatorID string, organizationID string) (*model.EntityProperty, error) {
	entity, err := s.entityService.Get(ctx, entityID, organizationID)

	if err != nil {
		return nil, err
	}

	property, err := s.propertyService.Get(ctx, propertyID, entity.ModelID, organizationID)

	if err != nil {
		return nil, err
	}

	if property.DataSchema != nil {
		err = property.DataSchema.Validate(request.Value)

		if err != nil {
			return nil, err
		}
	}

	entityProperty := &model.EntityProperty{
		EntityID:       entityID,
		PropertyID:     propertyID,
		Value:          request.Value,
		Rule:           request.Rule,
		CreatorType:    model.ActorTypeUser,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(entityProperty).CreateWithCtx(ctx, entityProperty)

	if err != nil {
		return nil, err
	}

	return entityProperty, nil
}

func (s *EntityPropertyService) List(ctx context.Context, entityID string, propertyID string, organizationID string, page int64, size int64) ([]*model.EntityProperty, error) {
	entityProperties := make([]*model.EntityProperty, 0)

	err := mgm.Coll(&model.EntityProperty{}).SimpleFindWithCtx(ctx, &entityProperties, bson.M{
		"entity_id":       entityID,
		"property_id":     propertyID,
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return entityProperties, nil
}

func (s *EntityPropertyService) Get(ctx context.Context, entityPropertyID string, entityID string, propertyID string, organizationID string) (*model.EntityProperty, error) {
	id, err := primitive.ObjectIDFromHex(entityPropertyID)

	if err != nil {
		return nil, err
	}

	entityProperty := &model.EntityProperty{}

	err = mgm.Coll(entityProperty).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"entity_id":       entityID,
		"property_id":     propertyID,
		"organization_id": organizationID,
	}, entityProperty)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("entity property not found")
	}

	if err != nil {
		return nil, err
	}

	return entityProperty, nil
}

func (s *EntityPropertyService) Update(ctx context.Context, entityPropertyID string, request *request.EntityPropertyRequest, entityID string, propertyID string, organizationID string) (*model.EntityProperty, error) {
	entityProperty, err := s.Get(ctx, entityPropertyID, entityID, propertyID, organizationID)

	if err != nil {
		return nil, err
	}

	property, err := s.propertyService.Get(ctx, propertyID, entityID, organizationID)

	if err != nil {
		return nil, err
	}

	if property.DataSchema != nil {
		err = property.DataSchema.Validate(request.Value)

		if err != nil {
			return nil, err
		}
	}

	entityProperty.Value = request.Value

	err = mgm.Coll(entityProperty).UpdateWithCtx(ctx, entityProperty)

	if err != nil {
		return nil, err
	}

	return entityProperty, nil
}

func (s *EntityPropertyService) Delete(ctx context.Context, entityPropertyID string, entityID string, propertyID string, organizationID string) (*model.EntityProperty, error) {
	entityProperty, err := s.Get(ctx, entityPropertyID, entityID, propertyID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(entityProperty).DeleteWithCtx(ctx, entityProperty)

	if err != nil {
		return nil, err
	}

	return entityProperty, nil
}
