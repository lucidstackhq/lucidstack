package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"-" bson:"password"`
	Admin            bool   `json:"admin" bson:"admin"`
	CreatorID        string `json:"creator_id" bson:"creator_id"`
	OrganizationID   string `json:"organization_id" bson:"organization_id"`
}
