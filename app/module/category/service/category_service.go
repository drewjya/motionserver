package service

import (
	"motionserver/app/module/category/repository"
	"motionserver/app/module/category/request"
	"motionserver/app/module/category/response"
	"motionserver/utils/paginator"
)

type categoryService struct {
	Repo repository.CategoryRepository
}

type CategoryService interface {
	All(req request.CategoriesRequest) (categories []*response.Categories, paging paginator.Pagination, err error)
	Store(req request.CategoryRequest) (err error)
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		Repo: repo,
	}
}

func (_i *categoryService) All(req request.CategoriesRequest) (categories []*response.Categories, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetCategories(req)
	if err != nil {
		return
	}
	for _, v := range results {
		categories = append(categories, response.FromDomain(v))
	}
	return

}

func (_i *categoryService) Store(req request.CategoryRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}
