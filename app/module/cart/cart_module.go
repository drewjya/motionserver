package cart

import (
	"motionserver/app/middleware"
	"motionserver/app/module/cart/controller"
	"motionserver/app/module/cart/repository"
	"motionserver/app/module/cart/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type CartRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewCartModule = fx.Options(
	fx.Provide(repository.NewCartRepository),
	fx.Provide(service.NewCartService),
	fx.Provide(controller.NewController),
	fx.Provide(NewCartRouter),
)

func NewCartRouter(fiber *fiber.App, controller *controller.Controller) *CartRouter {
	return &CartRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *CartRouter) RegisterCartRoutes() {
	cartController := _i.Controller.Cart
	_i.App.Route("/cart", func(router fiber.Router) {
		router.Get("", middleware.Protected(false), cartController.Index)
	})
}
