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

type ActionService struct {
	modelService *ModelService
}

func NewActionService(modelService *ModelService) *ActionService {
	return &ActionService{modelService: modelService}
}

func (a *ActionService) Create(ctx context.Context, modelID string, request *request.ActionCreationRequest, creatorID string, organizationID string) (*model.Action, error) {
	modelExists, err := a.modelService.Exists(ctx, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if !modelExists {
		return nil, fmt.Errorf("model not found")
	}

	nameExists, err := a.nameExists(ctx, request.Name, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("action %s already exists", request.Name)
	}

	action := &model.Action{
		ModelID:        modelID,
		Name:           request.Name,
		Description:    request.Description,
		InputSchema:    request.InputSchema,
		OutputSchema:   request.OutputSchema,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(action).CreateWithCtx(ctx, action)

	if err != nil {
		return nil, err
	}

	return action, nil
}

func (a *ActionService) List(ctx context.Context, modelID string, organizationID string, page int64, size int64) ([]*model.Action, error) {
	actions := make([]*model.Action, 0)

	err := mgm.Coll(&model.Action{}).SimpleFindWithCtx(ctx, &actions, bson.M{
		"model_id":        modelID,
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (a *ActionService) Get(ctx context.Context, actionID string, modelID string, organizationID string) (*model.Action, error) {
	id, err := primitive.ObjectIDFromHex(actionID)

	if err != nil {
		return nil, err
	}

	action := &model.Action{}

	err = mgm.Coll(action).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"model_id":        modelID,
		"organization_id": organizationID,
	}, action)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("action not found")
	}

	if err != nil {
		return nil, err
	}

	return action, nil
}

func (a *ActionService) Update(ctx context.Context, actionID string, request *request.ActionUpdateRequest, modelID string, organizationID string) (*model.Action, error) {
	action, err := a.Get(ctx, actionID, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		count, err := mgm.Coll(action).CountDocuments(ctx, bson.M{
			field.ID:          bson.M{operator.Ne: action.ID},
			"name":            request.Name,
			"model_id":        modelID,
			"organization_id": organizationID,
		})

		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf("action %s already exists", request.Name)
		}

		action.Name = request.Name
	}

	action.Description = request.Description

	err = mgm.Coll(action).UpdateWithCtx(ctx, action)

	if err != nil {
		return nil, err
	}

	return action, nil
}

func (a *ActionService) Delete(ctx context.Context, actionID string, modelID string, organizationID string) (*model.Action, error) {
	action, err := a.Get(ctx, actionID, modelID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(action).DeleteWithCtx(ctx, action)

	if err != nil {
		return nil, err
	}

	return action, nil
}

func (a *ActionService) nameExists(ctx context.Context, name string, modelID string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.Action{}).CountDocuments(ctx, bson.M{"name": name, "model_id": modelID, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
