package controller

import (
	"motionserver/app/module/category/request"
	"motionserver/app/module/category/service"
	"motionserver/utils/paginator"
	"motionserver/utils/response"

	"github.com/gofiber/fiber/v2"
)

type categoryController struct {
	categoryService service.CategoryService
}

// GetYoutube implements CategoryController.
func (_i *categoryController) GetYoutube(c *fiber.Ctx) error {
	yt, err := _i.categoryService.GetYoutube()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve youtube"),
		Data:     yt,
	})

}

// SetYoutube implements CategoryController.
func (_i *categoryController) SetYoutube(c *fiber.Ctx) error {
	err := _i.categoryService.SetYoutube()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success set youtube"),
	})
}

type CategoryController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error

	SetYoutube(c *fiber.Ctx) error
	GetYoutube(c *fiber.Ctx) error
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (_i *categoryController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.CategoriesRequest
	req.Pagination = paginate

	categories, paging, err := _i.categoryService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve categories"),
		Data:     categories,
		Meta:     paging,
	})

}

func (_i *categoryController) Store(c *fiber.Ctx) error {
	req := new(request.CategoryRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}
	err := _i.categoryService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Category Successfully created"),
	})
}
