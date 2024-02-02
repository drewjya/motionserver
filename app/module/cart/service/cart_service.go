package service

import (
	"context"
	"fmt"
	"motionserver/app/module/cart/repository"
	"motionserver/app/module/cart/request"
	"motionserver/app/module/cart/response"
	"motionserver/internal/bootstrap/minio"
	"motionserver/utils/paginator"
)

type cartService struct {
	Repo  repository.CartRepository
	Minio *minio.Minio
}

type CartService interface {
	All(req request.CartsRequest) (carts []*response.Cart, paging paginator.Pagination, err error)

	Store(req request.CartRequest) (err error)
	Update(id uint64, req request.UpdateCartRequest) (err error)
}

func NewCartService(repo repository.CartRepository, Minio *minio.Minio) CartService {
	return &cartService{
		Repo:  repo,
		Minio: Minio,
	}

}

func (_i *cartService) All(req request.CartsRequest) (carts []*response.Cart, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.FindCartByUserId(req)
	if err != nil {
		return
	}
	ctx := context.Background()
	for _, v := range results {
		fmt.Println(v)
		product := v.Product
		image := _i.Minio.GenerateLink(ctx, product.Image)
		product.Image = image
		v.Product = product
		carts = append(carts, response.FromDomain(v))
	}
	return
}

func (_i *cartService) Store(req request.CartRequest) (err error) {
	request := req.ToDomain()
	
	return _i.Repo.Create(request)

}
func (_i *cartService) Update(id uint64, req request.UpdateCartRequest) (err error) {
	cart, err := _i.Repo.FindOne(uint(id))
	if err != nil {
		return err
	}
	cart.Quantity = int32(req.Quantity)
	return _i.Repo.Update(uint(id), cart)

}