package model

import "github.com/kamva/mgm/v3"

type Environment struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" json:"name"`
	Description      string `bson:"description" json:"description"`
	CreatorID        string `bson:"creator_id" json:"creator_id"`
	OrganizationID   string `bson:"organization_id" json:"organization_id"`
}
