package product

import (
	"motionserver/app/middleware"
	"motionserver/app/module/product/controller"
	"motionserver/app/module/product/repository"
	"motionserver/app/module/product/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ProductRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewProductModule = fx.Options(
	fx.Provide(repository.NewProductRepository),
	fx.Provide(service.NewProductService),
	fx.Provide(controller.NewController),
	fx.Provide(NewProductRouter),
)

func NewProductRouter(fiber *fiber.App, controller *controller.Controller) *ProductRouter {
	return &ProductRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *ProductRouter) RegisterProductRoutes() {
	productController := _i.Controller.Product
	_i.App.Route("/product", func(router fiber.Router) {
		router.Get("", productController.Index)
		router.Post("", middleware.Protected(false), productController.Store)
	})
}
