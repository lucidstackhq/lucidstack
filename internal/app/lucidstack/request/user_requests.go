package request

type UserSignUpRequest struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	OrganizationName string `json:"organization_name" binding:"required"`
	BillingEmail     string `json:"billing_email" binding:"required"`
}

type UserTokenRequest struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	OrganizationName string `json:"organization_name" binding:"required"`
}

type UserPasswordChangeRequest struct {
	Password string `json:"password" binding:"required"`
}

type UserAdditionRequest struct {
	Username string `json:"username" binding:"required"`
	Admin    bool   `json:"admin"`
}

type UserAdminUpdateRequest struct {
	Admin bool `json:"admin"`
}
