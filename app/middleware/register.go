package middleware

import (
	"motionserver/utils/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Middleware struct {
	App *fiber.App
	Cfg *config.Config
}

func NewMiddleware(app *fiber.App, cfg *config.Config) *Middleware {
	return &Middleware{app, cfg}

}

func (m *Middleware) Register() {
	m.App.Use(logger.New())
	m.App.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "http://localhost:3000, http://localhost:3000, https://*.motionsportindonesia.id",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
}
