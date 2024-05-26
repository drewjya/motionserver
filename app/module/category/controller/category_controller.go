package controller

import (
	"fmt"
	"motionserver/app/module/category/request"
	"motionserver/app/module/category/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/paginator"
	"motionserver/utils/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type categoryController struct {
	categoryService service.CategoryService
}

// Delete implements CategoryController.
func (_i *categoryController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = _i.categoryService.Delete(uint(idVal))
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Category Successfully created"),
	})
}

// Update implements CategoryController.
func (_i *categoryController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	req := new(request.CategoryRequest)
	file, errNew := c.FormFile("image")

	var vale koderor.KodeError

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
	fmt.Println(req)

	err = _i.categoryService.Update(*req, uint(idVal))
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Category Successfully created"),
	})
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
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error

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
	fmt.Println(req)

	err = _i.categoryService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Category Successfully created"),
	})
}
