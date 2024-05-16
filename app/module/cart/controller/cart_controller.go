package controller

import (
	"motionserver/app/middleware"
	"motionserver/app/module/cart/request"
	"motionserver/app/module/cart/service"
	"motionserver/utils/paginator"
	"motionserver/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type cartController struct {
	cartService service.CartService
}

// Delete implements CartController.
func (_i *cartController) Delete(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return err
	}
	err = _i.cartService.Delete(id)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     nil,
		Messages: response.RootMessage("success delete cart"),
		Code:     fiber.StatusCreated,
	})
}

// Update implements CartController.
func (_i *cartController) Update(c *fiber.Ctx) error {
	var req request.UpdateCartRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = _i.cartService.Update(id, req)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     nil,
		Messages: response.RootMessage("success update cart"),
		Code:     fiber.StatusCreated,
	})
}

type CartController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewCartController(cartService service.CartService) CartController {
	return &cartController{
		cartService: cartService,
	}
}

func (_i *cartController) Store(c *fiber.Ctx) error {
	var req request.CartRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	jwt := c.Locals("token").(*middleware.JWTClaims)
	id, err := strconv.ParseUint(jwt.ID, 10, 64)
	if err != nil {
		return err
	}
	req.UserId = uint(id)

	err = _i.cartService.Store(req)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     nil,
		Messages: response.RootMessage("success create cart"),
		Code:     fiber.StatusCreated,
	})
}

func (_i *cartController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.CartsRequest
	req.Pagination = paginate
	jwt := c.Locals("token").(*middleware.JWTClaims)
	id, err := strconv.ParseUint(jwt.ID, 10, 64)
	if err != nil {
		return err
	}
	req.UserId = uint(id)

	galleries, paging, err := _i.cartService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve galleries"),
		Data:     galleries,
		Meta:     paging,
	})
}
