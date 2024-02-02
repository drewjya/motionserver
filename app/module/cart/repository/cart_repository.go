package repository

import (
	"fmt"
	"motionserver/app/database/schema"
	"motionserver/app/module/cart/request"
	"motionserver/internal/bootstrap/database"
	"motionserver/utils/paginator"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *database.Database
}

type CartRepository interface {
	FindCartByUserId(req request.CartsRequest) (carts []*schema.Cart, paging paginator.Pagination, err error)
	FindOne(id uint) (cart *schema.Cart, err error)
	Create(cart *schema.Cart) (err error)
	Update(id uint, cart *schema.Cart) (err error)
}

func NewCartRepository(db *database.Database) CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (_i *cartRepository) FindCartByUserId(req request.CartsRequest) (carts []*schema.Cart, paging paginator.Pagination, err error) {
	account := schema.Account{}
	err = _i.DB.DB.Where("user_id = ?", req.UserId).First(&account).Error
	if err != nil {
		return
	}
	var count int64

	query := _i.DB.DB.Model(&schema.Cart{}).Preload("Product.Categories").Preload("Product").Where("account_id = ?", account.ID)
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&carts).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	paging = *req.Pagination

	return

}

func (_i *cartRepository) Create(cart *schema.Cart) (err error) {
	var cval *schema.Cart
	_i.DB.DB.Where(schema.Cart{
		AccountID: cart.AccountID,
		ProductID: cart.ProductID,
	}).First(&cart)
	if cval != nil {
		cart.Quantity += cval.Quantity
		return _i.DB.DB.Save(&cart).Error
	}
	return _i.DB.DB.Save(&cart).Error
}

func (_i *cartRepository) Update(id uint, cart *schema.Cart) (err error) {
	return _i.DB.DB.Model(&schema.Cart{}).
		Where(&schema.Cart{
			Model: gorm.Model{
				ID: id,
			},
		}).
		Updates(cart).Error
}

func (_i *cartRepository) FindOne(id uint) (cart *schema.Cart, err error) {
	if err = _i.DB.DB.First(&cart, id).Error; err != nil {
		return
	}
	return
}
