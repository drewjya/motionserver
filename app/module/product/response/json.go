package response

import (
	"motionserver/app/database/schema"
	"motionserver/app/module/category/response"
)

type Product struct {
	ID          uint64                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Price       uint64                `json:"price"`
	Stock       uint                  `json:"stock"`
	Image       string                `json:"image"`
	SerialCode  string                `json:"serial_code"`
	Categories  []response.Categories `json:"categories"`
}

func FromDomain(domain *schema.Product, image string) (product *Product) {
	if domain == nil {
		return nil
	}
	var categories []response.Categories
	if domain.Categories != nil {
		for _, v := range domain.Categories {
			categories = append(categories, *response.FromDomain(&v))
		}

	}
	return &Product{
		ID:          uint64(domain.Model.ID),
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		Stock:       domain.Stock,
		Image:       image,
		SerialCode:  domain.SerialCode,
		Categories:  categories,
	}
}
