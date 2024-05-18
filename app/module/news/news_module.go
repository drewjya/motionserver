package news

import (
	"motionserver/app/middleware"
	"motionserver/app/module/news/controller"
	"motionserver/app/module/news/repository"
	"motionserver/app/module/news/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type NewsRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewNewsModule = fx.Options(
	fx.Provide(repository.NewNewsRepository),
	fx.Provide(service.NewNewsService),
	fx.Provide(controller.NewController),
	fx.Provide(NewNewsRouter),
)

func NewNewsRouter(fiber *fiber.App, controller *controller.Controller) *NewsRouter {
	return &NewsRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *NewsRouter) RegisterNewsRoutes() {
	newsController := _i.Controller.News
	_i.App.Route("/news", func(router fiber.Router) {
		router.Get("", newsController.Index)
		router.Get("/id", newsController.FindOne)

		router.Post("", middleware.Protected(false), newsController.Store)
		router.Put("/:id", middleware.Protected(false), newsController.Update)
		router.Delete("/:id", middleware.Protected(false), newsController.Delete)
	})
}
