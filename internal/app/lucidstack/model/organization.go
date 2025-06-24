package model

import "github.com/kamva/mgm/v3"

type Organization struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" bson:"name"`
	BillingEmail     string `bson:"billing_email" bson:"billing_email"`
	CreatorID        string `json:"creator_id" bson:"creator_id"`
}
