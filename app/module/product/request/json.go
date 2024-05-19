package request

import (
	"mime/multipart"
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"

	"gorm.io/gorm"
)

type ProductsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

type PromotionProductRequest struct {
	ProductId uint64 `form:"product_id" json:"product_id" validate:"required"`
	Image     string
	File      *multipart.FileHeader
}

type ProductRequest struct {
	Name        string `form:"name" json:"name" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
	Price       uint64 `form:"price" json:"price" validate:"required"`
	Stock       uint   `form:"stock" json:"stock" validate:"required"`
	SerialCode  string `form:"serialCode" json:"serialCode" validate:"required"`
	Image       string
	Categories  []uint64
	File        *multipart.FileHeader
}

func (req *ProductRequest) ToDomain() *schema.Product {
	res := new(schema.Product)
	res.Name = req.Name
	res.Description = req.Description
	res.Price = req.Price
	res.Stock = req.Stock
	res.SerialCode = req.SerialCode
	res.Image = req.Image
	var categories []schema.Category

	if req.Categories != nil {
		for _, v := range req.Categories {
			categories = append(categories, schema.Category{
				Model: gorm.Model{
					ID: uint(v),
				},
			})
		}
		res.Categories = categories
	}
	return res

}

func (req *PromotionProductRequest) ToDomain() *schema.PromotionProduct {
	res := new(schema.PromotionProduct)

	res.Product = schema.Product{
		Model: gorm.Model{
			ID: uint(req.ProductId),
		},
	}

	return res

}
