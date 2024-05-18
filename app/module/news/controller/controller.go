package controller

import "motionserver/app/module/news/service"

type Controller struct {
	News NewsController
}

func NewController(cartServce service.NewsService) *Controller {
	return &Controller{
		News: NewNewsController(cartServce),
	}
}
