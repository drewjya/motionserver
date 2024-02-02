package request

import (
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"
)

// [Your request handling logic here]

type CartsRequest struct {
	UserId     uint
	Pagination *paginator.Pagination `json:"pagination"`
}

type CartRequest struct {
	AccountId uint
	ProductId uint   `json:"productId"`
	Quantity  uint32 `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity uint32 `json:"quantity"`
}

func (req *CartRequest) ToDomain() (r *schema.Cart) {
	res := new(schema.Cart)
	res.ProductID = req.ProductId
	res.Quantity = int32(req.Quantity)
	res.AccountID = req.AccountId
	return res
}
