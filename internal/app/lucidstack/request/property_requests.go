package request

import "github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"

type PropertyCreationRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	DataSchema  *model.DataSchema `json:"data_schema" bson:"data_schema"`
}

type PropertyUpdateRequest struct {
	Description string `json:"description"`
}
