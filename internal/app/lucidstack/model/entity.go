package model

import "github.com/kamva/mgm/v3"

type Entity struct {
	mgm.DefaultModel `bson:",inline"`
	ModelID          string   `json:"model_id" bson:"model_id"`
	EnvironmentID    string   `bson:"environment_id" json:"environment_id"`
	ParentIDs        []string `bson:"parent_ids" json:"parent_ids"`
	Name             string   `bson:"name" json:"name"`
	Description      string   `bson:"description" json:"description"`
	CreatorID        string   `bson:"creator_id" json:"creator_id"`
	OrganizationID   string   `bson:"organization_id" json:"organization_id"`
}
