package request

type AppCreationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type AppUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
