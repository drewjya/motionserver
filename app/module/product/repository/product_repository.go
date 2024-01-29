package repository

import (
	"motionserver/app/database/schema"
	"motionserver/app/module/product/request"
	"motionserver/internal/bootstrap/database"
	"motionserver/utils/paginator"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *database.Database
}

type ProductRepository interface {
	GetProducts(req request.ProductsRequest) (products []*schema.Product, paging paginator.Pagination, err error)
	Create(product *schema.Product) (err error)
	Update(id uint64, product *schema.Product) (err error)
	Delete(id uint64) (err error)
}

func NewProductRepository(db *database.Database) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (_i *productRepository) GetProducts(req request.ProductsRequest) (products []*schema.Product, paging paginator.Pagination, err error) {
	var count int64
	query := _i.DB.DB.Model(&schema.Product{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&products).Error
	if err != nil {
		return
	}
	paging = *req.Pagination

	return
}

func (_i *productRepository) Create(product *schema.Product) (err error) {
	return _i.DB.DB.Create(&product).Error
}

func (_i *productRepository) Update(id uint64, product *schema.Product) (err error) {
	return _i.DB.DB.Model(&schema.Product{}).
		Where(&schema.Product{
			Model: gorm.Model{
				ID: uint(id),
			},
		}).
		Updates(product).Error
}

func (_i *productRepository) Delete(id uint64) (err error) {
	return _i.DB.DB.Delete(&schema.Product{}, id).Error
}