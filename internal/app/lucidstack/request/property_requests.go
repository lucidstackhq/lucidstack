package request

import "github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"

type PropertyCreationRequest struct {
	Name         string            `json:"name" binding:"required"`
	Description  string            `json:"description"`
	DataSchema   *model.DataSchema `json:"data_schema"`
	DefaultValue interface{}       `json:"default_value"`
}

type PropertyUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PropertyDefaultValueRequest struct {
	DefaultValue interface{} `json:"default_value"`
}
