package controller

import "motionserver/app/module/compro/service"

type Controller struct {
	Compro ComproController
}

func NewController(galleryServce service.ComproService) *Controller {
	return &Controller{
		Compro: NewComproController(galleryServce),
	}
}
