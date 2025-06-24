package auth

type AuthenticatedUser struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Admin          bool   `json:"admin"`
}
