package model

import "github.com/kamva/mgm/v3"

type Action struct {
	mgm.DefaultModel `bson:",inline"`
	ModelID          string      `json:"model_id" bson:"model_id"`
	Name             string      `json:"name" bson:"name"`
	Description      string      `json:"description" bson:"description"`
	InputSchema      *DataSchema `json:"input_schema" bson:"input_schema"`
	OutputSchema     *DataSchema `json:"output_schema" bson:"output_schema"`
	CreatorID        string      `json:"creator_id" bson:"creator_id"`
	OrganizationID   string      `json:"organization_id" bson:"organization_id"`
}
