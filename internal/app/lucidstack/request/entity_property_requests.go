package request

import "github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"

type EntityPropertyRequest struct {
	Value interface{} `json:"value" binding:"required"`
	Rule  *model.Rule `json:"rule"`
}
