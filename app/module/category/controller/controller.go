package controller

import "motionserver/app/module/category/service"

type Controller struct {
	Cateogry CategoryController
}

func NewController(categoryService service.CategoryService) *Controller {
	return &Controller{
		Cateogry: NewCategoryController(categoryService),
	}
}
