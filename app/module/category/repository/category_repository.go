package repository

import (
	"motionserver/app/database/schema"
	"motionserver/app/module/category/request"
	"motionserver/internal/bootstrap/database"
	"motionserver/utils/paginator"
)

type categoryRepository struct {
	DB *database.Database
}

type CategoryRepository interface {
	GetCategories(req request.CategoriesRequest) (categories []*schema.Category, paging paginator.Pagination, err error)
	Create(category *schema.Category) (err error)
}

func NewCategoryRepository(db *database.Database) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (_i *categoryRepository) GetCategories(req request.CategoriesRequest) (categories []*schema.Category, paging paginator.Pagination, err error) {
	// if err := _i.DB.DB.Find(&categories).Error; err != nil {
	// 	return nil, err
	// }
	// return categories, nil

	var count int64
	query := _i.DB.DB.Model(&schema.Category{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&categories).Error
	if err != nil {
		return
	}
	paging = *req.Pagination

	return
}

func (_i *categoryRepository) Create(category *schema.Category) (err error) {
	return _i.DB.DB.Create(&category).Error
}

