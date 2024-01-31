package controller

import (
	"motionserver/app/module/product/request"
	"motionserver/app/module/product/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/paginator"
	"motionserver/utils/response"

	"github.com/gofiber/fiber/v2"
)

type productController struct {
	productService service.ProductService
}

type ProductController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (_i *productController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.ProductsRequest
	req.Pagination = paginate

	products, paging, err := _i.productService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve products"),
		Data:     products,
		Meta:     paging,
	})
}

func (_i *productController) Store(c *fiber.Ctx) error {
	req := new(request.ProductRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}
	file, err := c.FormFile("image")
	if err != nil {
		vale := koderor.New("image", err.Error())
		return vale
	}
	req.File = file

	err = _i.productService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Product Successfully created"),
	})
}
