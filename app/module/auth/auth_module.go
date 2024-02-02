package auth

import (
	"motionserver/app/middleware"
	"motionserver/app/module/auth/controller"
	"motionserver/app/module/auth/repository"
	"motionserver/app/module/auth/service"

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
		auth.Get("/profile", middleware.Protected(false), authController.Profile)
		auth.Get("/refresh", middleware.Protected(true), authController.Refresh)
		auth.Post("/login", authController.Login)
		auth.Post("/register", authController.Register)
	})
}
