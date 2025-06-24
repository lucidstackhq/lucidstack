package response

import "github.com/lucidstackhq/lucidstack/internal/app/lucidstack/model"

type UserTokenResponse struct {
	Token string `json:"token"`
}

type UserPasswordResponse struct {
	Password string      `json:"password"`
	User     *model.User `json:"user"`
}
