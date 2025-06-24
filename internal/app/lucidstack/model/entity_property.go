package model

import "github.com/kamva/mgm/v3"

type EntityProperty struct {
	mgm.DefaultModel `bson:",inline"`
	EntityID         string      `bson:"entity_id" json:"entity_id"`
	Value            interface{} `bson:"value" json:"value"`
	Rule             *Rule       `bson:"rule" json:"rule"`
	CreatorID        string      `bson:"creator_id" json:"creator_id"`
	OrganizationID   string      `bson:"organization_id" json:"organization_id"`
}
