package service

import (
	"context"
	"motionserver/app/module/product/repository"
	"motionserver/app/module/product/request"
	"motionserver/app/module/product/response"
	"motionserver/internal/bootstrap/minio"
	"motionserver/utils/paginator"
)

type productService struct {
	Repo  repository.ProductRepository
	Minio *minio.Minio
}

type ProductService interface {
	All(req request.ProductsRequest) (categories []*response.Product, paging paginator.Pagination, err error)
	Show(id uint64) (product *response.Product, err error)
	Store(req request.ProductRequest) (err error)
}

func NewProductService(repo repository.ProductRepository, minio *minio.Minio) ProductService {
	return &productService{
		Repo:  repo,
		Minio: minio,
	}
}

func (_i *productService) All(req request.ProductsRequest) (products []*response.Product, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetProducts(req)
	if err != nil {
		return
	}
	ctx := context.Background()
	for _, v := range results {
		img := _i.Minio.GenerateLink(ctx, v.Image)
		products = append(products, response.FromDomain(v, img))
	}
	return
}

func (_i *productService) Show(id uint64) (product *response.Product, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return
	}
	ctx := context.Background()
	img := _i.Minio.GenerateLink(ctx, result.Image)
	product = response.FromDomain(result, img)
	return
}

func (_i *productService) Store(req request.ProductRequest) (err error) {
	if req.File != nil {
		ctx := context.Background()
		val, err := _i.Minio.UploadFile(ctx, *req.File)
		if err != nil {
			return err
		}
		req.Image = *val
	}
	return _i.Repo.Create(req.ToDomain())
}
