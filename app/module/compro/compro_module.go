package compro

import (
	"motionserver/app/middleware"
	"motionserver/app/module/compro/controller"
	"motionserver/app/module/compro/repository"
	"motionserver/app/module/compro/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ComproRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewComproModule = fx.Options(
	fx.Provide(repository.NewComproRepository),
	fx.Provide(service.NewComproService),
	fx.Provide(controller.NewController),
	fx.Provide(NewComproRouter),
)

func NewComproRouter(fiber *fiber.App, controller *controller.Controller) *ComproRouter {
	return &ComproRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *ComproRouter) RegisterComproRoutes() {
	comproController := _i.Controller.Compro
	_i.App.Route("/compro", func(router fiber.Router) {
		router.Get("", comproController.Index)
		router.Post("", middleware.Protected(false), comproController.Store)
	})
}
