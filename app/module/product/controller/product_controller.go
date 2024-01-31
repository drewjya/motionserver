package controller

import (
	"motionserver/app/module/product/request"
	"motionserver/app/module/product/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/paginator"
	"motionserver/utils/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	productService service.ProductService
}

type ProductController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
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

func (_i *productController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}
	product, err := _i.productService.Show(id)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve product"), Data: product})

}

func (_i *productController) Store(c *fiber.Ctx) error {
	req := new(request.ProductRequest)
	file, errNew := c.FormFile("image")
	var vale koderor.KodeError
	var err error
	if errNew != nil {
		errVall := errNew.Error()
		if errVall == "there is no uploaded file associated with the given key" {
			errVall = "Image field is required"
		}

		vale = koderor.New("image", errVall)
	}
	err = response.ParseAndValidate(c, req)
	if err != nil || vale != nil {
		if err != nil && vale == nil {
			val := err.(validator.ValidationErrors)
			return koderor.NewErrors(&val, nil)
		}
		if err == nil && vale != nil {
			return koderor.NewErrors(nil, vale.(*koderor.ErrorKode))
		}
		val := err.(validator.ValidationErrors)
		return koderor.NewErrors(&val, vale.(*koderor.ErrorKode))
	}
	req.File = file
	req.File = file

	err = _i.productService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Product Successfully created"),
	})
}
