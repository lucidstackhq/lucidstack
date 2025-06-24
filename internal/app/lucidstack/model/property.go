package model

import "github.com/kamva/mgm/v3"

type Property struct {
	mgm.DefaultModel `bson:",inline"`
	ModelID          string      `json:"model_id" bson:"model_id"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	DataSchema       *DataSchema `json:"data_schema" bson:"data_schema"`
	CreatorID        string      `json:"creator_id" bson:"creator_id"`
	OrganizationID   string      `json:"organization_id" bson:"organization_id"`
}
