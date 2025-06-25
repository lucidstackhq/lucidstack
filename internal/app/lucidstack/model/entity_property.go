package model

import "github.com/kamva/mgm/v3"

type EntityProperty struct {
	mgm.DefaultModel `bson:",inline"`
	EntityID         string      `bson:"entity_id" json:"entity_id"`
	PropertyID       string      `bson:"property_id" json:"property_id"`
	Value            interface{} `bson:"value" json:"value"`
	Rule             *Rule       `bson:"rule" json:"rule"`
	CreatorType      ActorType   `bson:"creator_type" json:"creator_type"`
	CreatorID        string      `bson:"creator_id" json:"creator_id"`
	OrganizationID   string      `bson:"organization_id" json:"organization_id"`
}
