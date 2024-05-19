package repository

import (
	"motionserver/app/database/schema"
	"motionserver/app/module/category/request"
	"motionserver/internal/bootstrap/database"
	"motionserver/utils/paginator"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *database.Database
}

// DeleteYoutube implements CategoryRepository.
func (_i *categoryRepository) DeleteYoutube() (err error) {
	return _i.DB.DB.Exec("DELETE FROM youtubes").Error
}

type CategoryRepository interface {
	GetCategories(req request.CategoriesRequest) (categories []*schema.Category, paging paginator.Pagination, err error)
	Create(category *schema.Category) (err error)
	GetYoutube() (youtube *schema.Youtube, err error)
	SetYoutube(youtube *schema.Youtube) (err error)
	DeleteYoutube() (err error)
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

func (_i *categoryRepository) GetYoutube() (youtube *schema.Youtube, err error) {
	if err := _i.DB.DB.First(&youtube).Error; err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
		return nil, nil
	}
	return youtube, nil
}
func (_i *categoryRepository) SetYoutube(youtube *schema.Youtube) (err error) {
	yt, err := _i.GetYoutube()
	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
		return _i.DB.DB.Model(&schema.Youtube{}).
			Where(&schema.Youtube{
				Model: gorm.Model{
					ID: yt.ID,
				},
			}).
			Updates(youtube).Error

	}
	return _i.DB.DB.Create(&youtube).Error
}
