package controller

import (
	"template_clean_architecture/app/module/auth/request"
	"template_clean_architecture/app/module/auth/service"
	"template_clean_architecture/utils/response"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (_i *authController) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}
	res, err := _i.authService.Login(*req)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.RootMessage("Login success"),

		Code: fiber.StatusOK,
	})
}

func (_i *authController) Register(c *fiber.Ctx) error {
	req := new(request.RegisterRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.Register(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.RootMessage("Register success"),
		Code:     fiber.StatusOK,
	})
}
