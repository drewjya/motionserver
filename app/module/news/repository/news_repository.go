package repository

import (
	"fmt"
	"motionserver/app/database/schema"
	"motionserver/app/module/news/request"
	"motionserver/internal/bootstrap/database"
	"motionserver/utils/paginator"

	"gorm.io/gorm"
)

type newsRepository struct {
	DB *database.Database
}

// DeleteNews implements NewsRepository.

type NewsRepository interface {
	FindNews(req request.NewssRequest) (newss []*schema.News, paging paginator.Pagination, err error)
	FindOne(id uint) (news *schema.News, err error)
	Create(news *schema.News) (err error)
	Update(id uint, news *schema.News) (err error)
	DeleteNews(id uint) (err error)
}

func NewNewsRepository(db *database.Database) NewsRepository {
	return &newsRepository{
		DB: db,
	}
}

func (_i *newsRepository) DeleteNews(id uint) (err error) {
	return _i.DB.DB.Unscoped().Delete(&schema.News{}, id).Error
}
func (_i *newsRepository) FindNews(req request.NewssRequest) (newss []*schema.News, paging paginator.Pagination, err error) {

	var count int64

	query := _i.DB.DB.Model(&schema.News{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Order("tanggal").Find(&newss).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	paging = *req.Pagination

	return

}

func (_i *newsRepository) Create(news *schema.News) (err error) {
	return _i.DB.DB.Create(&news).Error
}

func (_i *newsRepository) Update(id uint, news *schema.News) (err error) {
	return _i.DB.DB.Model(&schema.News{}).
		Where(&schema.News{
			Model: gorm.Model{
				ID: id,
			},
		}).
		Updates(news).Error
}

func (_i *newsRepository) FindOne(id uint) (news *schema.News, err error) {
	if err = _i.DB.DB.First(&news, id).Error; err != nil {
		return nil, err
	}
	return news, nil
}
