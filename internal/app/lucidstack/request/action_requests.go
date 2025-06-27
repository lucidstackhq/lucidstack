package request

import "github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"

type ActionCreationRequest struct {
	Name         string            `json:"name" binding:"required"`
	Description  string            `json:"description"`
	InputSchema  *model.DataSchema `json:"input_schema"`
	OutputSchema *model.DataSchema `json:"output_schema"`
}

type ActionUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
