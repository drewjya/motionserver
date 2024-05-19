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

// CreatePromotionProduct implements ProductRepository.
func (_i *productRepository) CreatePromotionProduct(product *schema.PromotionProduct) (err error) {
	return _i.DB.DB.Create(&product).Error
}

// DeletePromotionProduct implements ProductRepository.
func (_i *productRepository) DeletePromotionProduct(id uint64) (err error) {
	return _i.DB.DB.Delete(&schema.PromotionProduct{}, id).Error
}

// GetAllPromotionProduct implements ProductRepository.
func (_i *productRepository) GetAllPromotionProduct() (products []*schema.PromotionProduct, err error) {
	query := _i.DB.DB.Model(&schema.PromotionProduct{}).Preload("Product")
	err = query.Find(&products).Error
	if err != nil {
		return
	}

	return
}

// GetPromotionProduct implements ProductRepository.
func (_i *productRepository) GetPromotionProduct(id uint64) (product *schema.PromotionProduct, err error) {
	if err := _i.DB.DB.Model(&schema.PromotionProduct{}).Preload("Product").First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (_i *productRepository) GetPromotionProductByProduct(id uint64) (product *schema.PromotionProduct, err error) {
	if err := _i.DB.DB.Where("ProductID = ?", id).Preload("Product").First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

type ProductRepository interface {
	GetProducts(req request.ProductsRequest) (products []*schema.Product, paging paginator.Pagination, err error)
	FindOne(id uint64) (product *schema.Product, err error)
	Create(product *schema.Product) (err error)
	Update(id uint64, product *schema.Product) (err error)
	Delete(id uint64) (err error)

	GetAllPromotionProduct() (products []*schema.PromotionProduct, err error)
	GetPromotionProduct(id uint64) (product *schema.PromotionProduct, err error)
	GetPromotionProductByProduct(id uint64) (product *schema.PromotionProduct, err error)
	CreatePromotionProduct(product *schema.PromotionProduct) (err error)
	DeletePromotionProduct(id uint64) (err error)
}

func NewProductRepository(db *database.Database) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (_i *productRepository) GetProducts(req request.ProductsRequest) (products []*schema.Product, paging paginator.Pagination, err error) {
	var count int64
	query := _i.DB.DB.Model(&schema.Product{}).Preload("Categories")
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

func (_i *productRepository) FindOne(id uint64) (product *schema.Product, err error) {
	if err := _i.DB.DB.Model(&schema.Product{}).Preload("Categories").First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
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
