package service

import (
	"motionserver/app/module/product/repository"
	"motionserver/app/module/product/request"
	"motionserver/app/module/product/response"
	"motionserver/utils/paginator"
)

type productService struct {
	Repo repository.ProductRepository
}

type ProductService interface {
	All(req request.ProductsRequest) (categories []*response.Product, paging paginator.Pagination, err error)
	Store(req request.ProductRequest) (err error)
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		Repo: repo,
	}

}

func (_i *productService) All(req request.ProductsRequest) (products []*response.Product, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetProducts(req)
	if err != nil {
		return
	}
	for _, v := range results {
		products = append(products, response.FromDomain(v))
	}
	return
}

func (_i *productService) Store(req request.ProductRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}
