package controller

import "motionserver/app/module/product/service"

type Controller struct {
	Product ProductController
}

func NewController(productServce service.ProductService) *Controller {
	return &Controller{
		Product: NewProductController(productServce),
	}
}
