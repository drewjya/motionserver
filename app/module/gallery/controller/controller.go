package controller
import "motionserver/app/module/gallery/service"

type Controller struct {
	Gallery GalleryController
}

func NewController(galleryServce service.GalleryService) *Controller {
	return &Controller{
		Gallery: NewGalleryController(galleryServce),
	}
}



