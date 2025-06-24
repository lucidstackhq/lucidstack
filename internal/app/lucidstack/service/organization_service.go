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
)

type OrganizationService struct {
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func (s *OrganizationService) Save(ctx context.Context, name string, billingEmail string, creatorID string) (*model.Organization, error) {

	organization := &model.Organization{
		Name:         name,
		BillingEmail: billingEmail,
		CreatorID:    creatorID,
	}

	err := mgm.Coll(organization).CreateWithCtx(ctx, organization)

	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) Get(ctx context.Context, organizationID string) (*model.Organization, error) {
	id, err := primitive.ObjectIDFromHex(organizationID)

	if err != nil {
		return nil, err
	}

	organization := &model.Organization{}

	err = mgm.Coll(organization).FirstWithCtx(ctx, bson.M{
		field.ID: id,
	}, organization)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("organization not found")
	}

	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) GetByName(ctx context.Context, name string) (*model.Organization, error) {
	organization := &model.Organization{}

	err := mgm.Coll(organization).FirstWithCtx(ctx, bson.M{
		"name": name,
	}, organization)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("organization %s not found", name)
	}

	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) Update(ctx context.Context, organizationID string, request *request.OrganizationUpdateRequest) (*model.Organization, error) {

	organization, err := s.Get(ctx, organizationID)

	if err != nil {
		return nil, err
	}

	organization.BillingEmail = request.BillingEmail

	err = mgm.Coll(organization).UpdateWithCtx(ctx, organization)

	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) NameExists(ctx context.Context, name string) (bool, error) {
	count, err := mgm.Coll(&model.Organization{}).CountDocuments(ctx, bson.M{"name": name})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
