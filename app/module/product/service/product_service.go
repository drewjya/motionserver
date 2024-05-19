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

// CreatePromotion implements ProductService.
func (_i *productService) CreatePromotion(req request.PromotionProductRequest) (err error) {
	oldProduct, err := _i.Repo.GetPromotionProductByProduct(req.ProductId)
	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}
	if oldProduct != nil {
		ctx := context.Background()
		err = _i.Minio.DeleteFile(ctx, oldProduct.Image)
		if err != nil {
			return
		}
		err = _i.Repo.DeletePromotionProduct(uint64(oldProduct.ID))
		if err != nil {
			return
		}

	}
	ctx := context.Background()
	if req.File != nil {
		val, err := _i.Minio.UploadFile(ctx, *req.File)
		if err != nil {
			return err
		}
		req.Image = *val
	}
	return _i.Repo.CreatePromotionProduct(req.ToDomain())
}

// DeletePromotion implements ProductService.
func (_i *productService) DeletePromotion(id uint64) (err error) {
	oldProduct, err := _i.Repo.GetPromotionProduct(id)
	if err != nil {
		return err
	}
	if oldProduct != nil {
		ctx := context.Background()
		err = _i.Minio.DeleteFile(ctx, oldProduct.Image)
		if err != nil {
			return
		}
		err = _i.Repo.DeletePromotionProduct(uint64(oldProduct.ID))
		if err != nil {
			return
		}

	}
	return
}

// GetAllPromotion implements ProductService.
func (_i *productService) GetAllPromotion() (products []*response.PromotionProduct, err error) {
	results, err := _i.Repo.GetAllPromotionProduct()
	if err != nil {
		return
	}
	ctx := context.Background()
	for _, v := range results {
		imgPromotion := _i.Minio.GenerateLink(ctx, v.Image)
		img := _i.Minio.GenerateLink(ctx, v.Product.Image)
		products = append(products, response.FromPromotionProduct(v, imgPromotion, img))
	}
	return
}

// GetPromotion implements ProductService.
func (_i *productService) GetPromotion(id uint64) (product *response.PromotionProduct, err error) {
	result, err := _i.Repo.GetPromotionProduct(id)
	if err != nil {
		return
	}
	ctx := context.Background()

	imgPromotion := _i.Minio.GenerateLink(ctx, result.Image)
	img := _i.Minio.GenerateLink(ctx, result.Product.Image)
	product = response.FromPromotionProduct(result, imgPromotion, img)

	return
}

type ProductService interface {
	All(req request.ProductsRequest) (categories []*response.Product, paging paginator.Pagination, err error)
	Show(id uint64) (product *response.Product, err error)
	Store(req request.ProductRequest) (err error)

	GetAllPromotion() (products []*response.PromotionProduct, err error)
	GetPromotion(id uint64) (product *response.PromotionProduct, err error)
	CreatePromotion(req request.PromotionProductRequest) (err error)
	DeletePromotion(id uint64) (err error)
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
