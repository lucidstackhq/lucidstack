package request

type ModelCreationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type ModelUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
