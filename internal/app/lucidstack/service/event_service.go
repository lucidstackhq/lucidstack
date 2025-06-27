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

type EventService struct {
	modelService *ModelService
}

func NewEventService(modelService *ModelService) *EventService {
	return &EventService{modelService: modelService}
}

func (s *EventService) Create(ctx context.Context, modelID string, request *request.EventCreationRequest, creatorID string, organizationID string) (*model.Event, error) {
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
		return nil, fmt.Errorf("event %s already exists", request.Name)
	}

	event := &model.Event{
		ModelID:        modelID,
		Name:           request.Name,
		Description:    request.Description,
		DataSchema:     request.DataSchema,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(event).CreateWithCtx(ctx, event)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) List(ctx context.Context, modelID string, organizationID string, page int64, size int64) ([]*model.Event, error) {
	events := make([]*model.Event, 0)

	err := mgm.Coll(&model.Event{}).SimpleFindWithCtx(ctx, &events, bson.M{
		"model_id":        modelID,
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *EventService) Get(ctx context.Context, eventID string, modelID string, organizationID string) (*model.Event, error) {
	id, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return nil, err
	}

	event := &model.Event{}

	err = mgm.Coll(event).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"model_id":        modelID,
		"organization_id": organizationID,
	}, event)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("event not found")
	}

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) Update(ctx context.Context, eventID string, request *request.EventUpdateRequest, modelID string, organizationID string) (*model.Event, error) {
	event, err := s.Get(ctx, eventID, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		count, err := mgm.Coll(event).CountDocuments(ctx, bson.M{
			field.ID: bson.M{
				operator.Ne:       bson.M{operator.Ne: event.ID},
				"name":            request.Name,
				"model_id":        modelID,
				"organization_id": organizationID,
			},
		})

		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf("event %s already exists", request.Name)
		}

		event.Name = request.Name
	}

	event.Description = request.Description

	err = mgm.Coll(event).UpdateWithCtx(ctx, event)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) Delete(ctx context.Context, eventID string, modelID string, organizationID string) (*model.Event, error) {
	event, err := s.Get(ctx, eventID, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(event).DeleteWithCtx(ctx, event)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) nameExists(ctx context.Context, name string, modelID string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.Event{}).CountDocuments(ctx, bson.M{
		"name":            name,
		"model_id":        modelID,
		"organization_id": organizationID,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
