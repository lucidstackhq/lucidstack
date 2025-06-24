package request

type EnvironmentCreationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type EnvironmentUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
