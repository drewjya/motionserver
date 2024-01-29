package controller

import "template_clean_architecture/app/module/auth/service"

type Controller struct {
	Auth AuthController
}

func NewController(
	authService service.AuthService,
) *Controller {
	return &Controller{
		Auth: NewAuthController(authService),
	}
}
