package controller

import (
	"motionserver/app/middleware"
	"motionserver/app/module/auth/request"
	"motionserver/app/module/auth/service"
	"motionserver/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) (err error)
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (_i *authController) Refresh(c *fiber.Ctx) (err error) {
	jwt := c.Locals("token").(*middleware.JWTClaims)
	id, err := strconv.ParseUint(jwt.ID, 10, 64)
	if err != nil {
		return
	}
	res, err := _i.authService.RefreshToken(id)
	if err != nil {
		return
	}
	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.RootMessage("Refresh token success"),
		Code:     fiber.StatusOK,
	})

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
		Code:     fiber.StatusOK,
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
