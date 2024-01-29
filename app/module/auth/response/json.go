package response

import (
	"template_clean_architecture/app/middleware"
)

type LoginResponse struct {
	Name      string                   `json:"name"`
	Email     string                   `json:"email"`
	UserId    uint64                   `json:"userId"`
	AccountId uint64                   `json:"accountId"`
	Token     middleware.TokenResponse `json:"token"`
}

type RegisterResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	UserId    uint64 `json:"userId"`
	AccountId uint64 `json:"accountId"`
}
