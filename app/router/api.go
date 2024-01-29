package router

import (
	"motionserver/app/module/auth"
	"motionserver/utils/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	App fiber.Router
	Cfg *config.Config

	AuthRouter *auth.AuthRouter
}

func NewRouter(
	fiber *fiber.App,
	cfg *config.Config,

	authRouter *auth.AuthRouter,
) *Router {
	return &Router{
		App: fiber,
		Cfg: cfg,

		AuthRouter: authRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Swagger Documentation
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// Register routes of modules
	r.AuthRouter.RegisterAuthRoutes()
}
