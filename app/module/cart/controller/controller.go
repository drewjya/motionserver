package controller
import "motionserver/app/module/cart/service"

type Controller struct {
	Cart CartController
}

func NewController(cartServce service.CartService) *Controller {
	return &Controller{
		Cart: NewCartController(cartServce),
	}
}



