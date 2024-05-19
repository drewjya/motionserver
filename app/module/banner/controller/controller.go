package controller

import "motionserver/app/module/banner/service"

type Controller struct {
	Banner BannerController
}

func NewController(bannerServce service.BannerService) *Controller {
	return &Controller{
		Banner: NewBannerController(bannerServce),
	}
}
