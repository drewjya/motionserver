package category

import (
	"motionserver/app/middleware"
	"motionserver/app/module/category/controller"
	"motionserver/app/module/category/repository"
	"motionserver/app/module/category/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type CategoryRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewCategoryModule = fx.Options(
	fx.Provide(repository.NewCategoryRepository),
	fx.Provide(service.NewCategoryService),
	fx.Provide(controller.NewController),
	fx.Provide(NewCategoryRouter),
)

func NewCategoryRouter(fiber *fiber.App, controller *controller.Controller) *CategoryRouter {
	return &CategoryRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *CategoryRouter) RegisterCategoryRoutes() {
	categoryController := _i.Controller.Cateogry
	_i.App.Route("/category", func(router fiber.Router) {
		router.Get("", categoryController.Index)

		router.Post("", middleware.Protected(false), categoryController.Store)
	})
	_i.App.Route("/youtube", func(router fiber.Router) {
		router.Get("", categoryController.GetYoutube)

		router.Post("", categoryController.SetYoutube)
	})
}
