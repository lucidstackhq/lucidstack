package request

import "github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"

type EventCreationRequest struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description"`
	DataSchema  *model.DataSchema `json:"data_schema"`
}

type EventUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
