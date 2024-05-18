package router

import (
	"motionserver/app/module/auth"
	"motionserver/app/module/cart"
	"motionserver/app/module/category"
	"motionserver/app/module/gallery"
	"motionserver/app/module/news"
	"motionserver/app/module/product"
	"motionserver/utils/config"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App            fiber.Router
	Cfg            *config.Config
	AuthRouter     *auth.AuthRouter
	CategoryRouter *category.CategoryRouter
	ProductRouter  *product.ProductRouter
	GalleryRouter  *gallery.GalleryRouter
	CartRouter     *cart.CartRouter
	NewsRouter     *news.NewsRouter
}

func NewRouter(
	fiber *fiber.App,
	cfg *config.Config,
	authRouter *auth.AuthRouter,
	categoryRouter *category.CategoryRouter,
	ProductRouter *product.ProductRouter,
	GalleryRouter *gallery.GalleryRouter,
	CartRouter *cart.CartRouter,
	NewsRouter *news.NewsRouter,
) *Router {
	return &Router{
		App:            fiber,
		Cfg:            cfg,
		AuthRouter:     authRouter,
		CategoryRouter: categoryRouter,
		ProductRouter:  ProductRouter,
		GalleryRouter:  GalleryRouter,
		CartRouter:     CartRouter,
		NewsRouter:     NewsRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Swagger Documentation

	// Register routes of modules
	r.AuthRouter.RegisterAuthRoutes()

	r.CategoryRouter.RegisterCategoryRoutes()

	r.ProductRouter.RegisterProductRoutes()

	r.GalleryRouter.RegisterGalleryRoutes()
	r.CartRouter.RegisterCartRoutes()
	r.NewsRouter.RegisterNewsRoutes()

}
