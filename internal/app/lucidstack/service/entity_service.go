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
	"time"
)

type EntityService struct {
	modelService       *ModelService
	environmentService *EnvironmentService
}

func NewEntityService(modelService *ModelService, environmentService *EnvironmentService) *EntityService {
	return &EntityService{modelService: modelService, environmentService: environmentService}
}

func (s *EntityService) Create(ctx context.Context, request *request.EntityCreationRequest, creatorID string, organizationID string) (*model.Entity, error) {
	modelExists, err := s.modelService.Exists(ctx, request.ModelID, organizationID)

	if err != nil {
		return nil, err
	}

	if !modelExists {
		return nil, fmt.Errorf("model not found")
	}

	environmentExists, err := s.environmentService.Exists(ctx, request.EnvironmentID, organizationID)

	if err != nil {
		return nil, err
	}

	if !environmentExists {
		return nil, fmt.Errorf("environment not found")
	}

	if request.ParentIDs != nil && len(request.ParentIDs) != 0 {
		allParentsExist, err := s.allExist(ctx, request.ParentIDs, request.EnvironmentID, organizationID)

		if err != nil {
			return nil, err
		}

		if !allParentsExist {
			return nil, fmt.Errorf("parent entities not found")
		}
	} else {
		request.ParentIDs = make([]string, 0)
	}

	nameExists, err := s.nameExists(ctx, request.Name, request.ModelID, request.EnvironmentID, organizationID)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("entity %s already exists", request.Name)
	}

	entity := &model.Entity{
		ModelID:        request.ModelID,
		EnvironmentID:  request.EnvironmentID,
		ParentIDs:      request.ParentIDs,
		Name:           request.Name,
		Description:    request.Description,
		CreatorID:      creatorID,
		OrganizationID: organizationID,
	}

	err = mgm.Coll(entity).CreateWithCtx(ctx, entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *EntityService) ListForModel(ctx context.Context, modelID string, environmentID string, organizationID string, page int64, size int64) ([]*model.Entity, error) {
	entities := make([]*model.Entity, 0)

	err := mgm.Coll(&model.Entity{}).SimpleFindWithCtx(ctx, &entities, bson.M{
		"model_id":        modelID,
		"environment_id":  environmentID,
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (s *EntityService) ListForParent(ctx context.Context, parentID string, environmentID string, organizationID string, page int64, size int64) ([]*model.Entity, error) {
	entities := make([]*model.Entity, 0)

	err := mgm.Coll(&model.Entity{}).SimpleFindWithCtx(ctx, &entities, bson.M{
		"parent_ids":      parentID,
		"environment_id":  environmentID,
		"organization_id": organizationID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (s *EntityService) Get(ctx context.Context, entityID string, organizationID string) (*model.Entity, error) {
	id, err := primitive.ObjectIDFromHex(entityID)

	if err != nil {
		return nil, err
	}

	entity := &model.Entity{}

	err = mgm.Coll(entity).FirstWithCtx(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	}, entity)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("entity not found")
	}

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *EntityService) Update(ctx context.Context, entityID string, request *request.EntityUpdateRequest, organizationID string) (*model.Entity, error) {
	entity, err := s.Get(ctx, entityID, organizationID)

	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		count, err := mgm.Coll(entity).CountDocuments(ctx, bson.M{
			field.ID:          bson.M{operator.Ne: entity.ID},
			"name":            request.Name,
			"environment_id":  entity.EnvironmentID,
			"model_id":        entity.ModelID,
			"organization_id": entity.OrganizationID,
		})

		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, fmt.Errorf("entity %s already exists", request.Name)
		}

		entity.Name = request.Name
	}

	entity.Description = request.Description

	err = mgm.Coll(entity).UpdateWithCtx(ctx, entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *EntityService) Delete(ctx context.Context, entityID string, organizationID string) (*model.Entity, error) {
	entity, err := s.Get(ctx, entityID, organizationID)

	if err != nil {
		return nil, err
	}

	err = mgm.Coll(entity).DeleteWithCtx(ctx, entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *EntityService) AddParent(ctx context.Context, entityID string, request *request.EntityParentRequest, organizationID string) error {
	entity, err := s.Get(ctx, entityID, organizationID)

	if err != nil {
		return err
	}

	parentExists, err := s.exists(ctx, request.ParentID, entity.EnvironmentID, organizationID)

	if err != nil {
		return err
	}

	if !parentExists {
		return fmt.Errorf("parent entity not found")
	}

	result, err := mgm.Coll(entity).UpdateOne(ctx, bson.M{
		field.ID:          entity.ID,
		"organization_id": organizationID,
	}, bson.M{
		operator.AddToSet: bson.M{
			"parent_ids": request.ParentID,
		},
		operator.Set: bson.M{
			"updated_at": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("entity not found")
	}

	return nil
}

func (s *EntityService) DeleteParent(ctx context.Context, entityID string, request *request.EntityParentRequest, organizationID string) error {
	id, err := primitive.ObjectIDFromHex(entityID)

	if err != nil {
		return err
	}

	result, err := mgm.Coll(&model.Entity{}).UpdateOne(ctx, bson.M{
		field.ID:          id,
		"organization_id": organizationID,
	}, bson.M{
		operator.Pull: bson.M{
			"parent_ids": request.ParentID,
		},
		operator.Set: bson.M{
			"updated_at": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("entity not found")
	}

	return nil
}

func (s *EntityService) ListParents(ctx context.Context, entityID string, organizationID string) ([]*model.Entity, error) {
	entity, err := s.Get(ctx, entityID, organizationID)

	if err != nil {
		return nil, err
	}

	if entity.ParentIDs == nil || len(entity.ParentIDs) == 0 {
		return []*model.Entity{}, nil
	}

	return s.getByIDs(ctx, entity.ParentIDs, organizationID)
}

func (s *EntityService) nameExists(ctx context.Context, name string, modelID string, environmentID string, organizationID string) (bool, error) {
	count, err := mgm.Coll(&model.Entity{}).CountDocuments(ctx, bson.M{"name": name, "model_id": modelID, "environment_id": environmentID, "organization_id": organizationID})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *EntityService) exists(ctx context.Context, entityID string, environmentID string, organizationID string) (bool, error) {
	id, err := primitive.ObjectIDFromHex(entityID)

	if err != nil {
		return false, err
	}

	count, err := mgm.Coll(&model.Entity{}).CountDocuments(ctx, bson.M{
		field.ID:          id,
		"environment_id":  environmentID,
		"organization_id": organizationID,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *EntityService) allExist(ctx context.Context, entityIDs []string, environmentID string, organizationID string) (bool, error) {
	ids := make([]primitive.ObjectID, len(entityIDs))

	for i, entityID := range entityIDs {
		id, err := primitive.ObjectIDFromHex(entityID)

		if err != nil {
			return false, err
		}

		ids[i] = id
	}

	count, err := mgm.Coll(&model.Entity{}).CountDocuments(ctx, bson.M{
		field.ID: bson.M{
			operator.In: ids,
		},
		"environment_id":  environmentID,
		"organization_id": organizationID,
	})

	if err != nil {
		return false, err
	}

	return int(count) == len(entityIDs), nil
}

func (s *EntityService) getByIDs(ctx context.Context, entityIDs []string, organizationID string) ([]*model.Entity, error) {
	ids := make([]primitive.ObjectID, len(entityIDs))

	for i, entityID := range entityIDs {
		id, err := primitive.ObjectIDFromHex(entityID)

		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	entities := make([]*model.Entity, 0)

	err := mgm.Coll(&model.Entity{}).SimpleFindWithCtx(ctx, &entities, bson.M{
		field.ID: bson.M{
			operator.In: ids,
		},
		"organization_id": organizationID,
	})

	if err != nil {
		return nil, err
	}

	return entities, nil
}
