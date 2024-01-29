package auth

import (
	"template_clean_architecture/app/module/auth/controller"
	"template_clean_architecture/app/module/auth/repository"
	"template_clean_architecture/app/module/auth/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AuthRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

var NewAuthModule = fx.Options(
	fx.Provide(repository.NewAuthRepository),
	fx.Provide(service.NewAuthService),
	fx.Provide(controller.NewController),
	fx.Provide(NewAuthRouter),
)

func NewAuthRouter(fiber *fiber.App, controller *controller.Controller) *AuthRouter {
	return &AuthRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (_i *AuthRouter) RegisterAuthRoutes() {

	authController := _i.Controller.Auth
	_i.App.Route("/auth", func(auth fiber.Router) {
		auth.Post("/login", authController.Login)
		auth.Post("/register", authController.Register)
	})
}
