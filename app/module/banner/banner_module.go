package banner

import (
	"motionserver/app/middleware"
	"motionserver/app/module/banner/controller"
	"motionserver/app/module/banner/repository"
	"motionserver/app/module/banner/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type BannerRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewBannerModule = fx.Options(
	fx.Provide(repository.NewBannerRepository),
	fx.Provide(service.NewBannerService),
	fx.Provide(controller.NewController),
	fx.Provide(NewBannerRouter),
)

func NewBannerRouter(fiber *fiber.App, controller *controller.Controller) *BannerRouter {
	return &BannerRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *BannerRouter) RegisterBannerRoutes() {
	bannerController := _i.Controller.Banner
	_i.App.Route("/banner", func(router fiber.Router) {
		router.Get("", bannerController.Index)
		router.Post("", middleware.Protected(false), bannerController.Store)
	})
}
