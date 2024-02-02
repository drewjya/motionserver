package gallery

import (
	"motionserver/app/middleware"
	"motionserver/app/module/gallery/controller"
	"motionserver/app/module/gallery/repository"
	"motionserver/app/module/gallery/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type GalleryRouter struct {
	App        *fiber.App
	Controller controller.Controller
}

var NewGalleryModule = fx.Options(
	fx.Provide(repository.NewGalleryRepository),
	fx.Provide(service.NewGalleryService),
	fx.Provide(controller.NewController),
	fx.Provide(NewGalleryRouter),
)

func NewGalleryRouter(fiber *fiber.App, controller *controller.Controller) *GalleryRouter {
	return &GalleryRouter{
		App:        fiber,
		Controller: *controller,
	}

}

func (_i *GalleryRouter) RegisterGalleryRoutes() {
	galleryController := _i.Controller.Gallery
	_i.App.Route("/gallery", func(router fiber.Router) {
		router.Get("", galleryController.Index)
		router.Get("/:id", galleryController.Show)
		router.Post("", middleware.Protected(false), galleryController.Store)
	})
}
