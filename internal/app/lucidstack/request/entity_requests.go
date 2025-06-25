package request

type EntityCreationRequest struct {
	ModelID       string   `json:"model_id" binding:"required"`
	EnvironmentID string   `json:"environment_id" binding:"required"`
	ParentIDs     []string `json:"parent_ids"`
	Name          string   `json:"name" binding:"required"`
	Description   string   `json:"description"`
}

type EntityUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EntityParentRequest struct {
	ParentID string `json:"parent_id" binding:"required"`
}
