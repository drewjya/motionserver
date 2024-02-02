package response

import (
	"motionserver/app/database/schema"
	"motionserver/app/module/product/response"
)

// [Your response handling logic here]

type Cart struct {
	ID       uint             `json:"id"`
	Quantity int32            `json:"quantity"`
	Product  response.Product `json:"product"`
}

func FromDomain(domain *schema.Cart) (cart *Cart) {
	if domain == nil {
		return nil
	}

	return &Cart{
		ID:       domain.ID,
		Quantity: domain.Quantity,
		Product:  *response.FromDomain(&domain.Product, domain.Product.Image),
	}
}
