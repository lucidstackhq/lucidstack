package request

type OrganizationUpdateRequest struct {
	BillingEmail string `json:"billing_email" binding:"required"`
}
