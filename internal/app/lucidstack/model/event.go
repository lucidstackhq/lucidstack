package model

import "github.com/kamva/mgm/v3"

type Event struct {
	mgm.DefaultModel `bson:",inline"`
	ModelID          string `json:"model_id" bson:"model_id"`
	Name             string `bson:"name" json:"name"`
	Description      string `bson:"description" json:"description"`
	DataSchema       string `bson:"data_schema" json:"data_schema"`
	CreatorID        string `bson:"creator_id" json:"creator_id"`
	OrganizationID   string `bson:"organization_id" json:"organization_id"`
}
