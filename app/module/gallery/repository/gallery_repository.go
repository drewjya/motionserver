package repository

import (
	"motionserver/app/database/schema"
	"motionserver/app/module/gallery/request"
	"motionserver/internal/bootstrap/database"
	"motionserver/utils/paginator"

	"gorm.io/gorm"
)

type galleryRepository struct {
	DB *database.Database
}

type GalleryRepository interface {
	GetProducts(req request.GalleriesRequest) (galleries []*schema.Gallery, paging paginator.Pagination, err error)
	Create(gallery *schema.Gallery) (err error)
	Update(id uint64, gallery *schema.Gallery) (err error)
	Delete(id uint64) (err error)
}

func NewGalleryRepository(db *database.Database) GalleryRepository {
	return &galleryRepository{
		DB: db,
	}
}

func (_i *galleryRepository) GetProducts(req request.GalleriesRequest) (galleries []*schema.Gallery, paging paginator.Pagination, err error) {
	var count int64
	query := _i.DB.DB.Model(&schema.Gallery{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&galleries).Error
	if err != nil {
		return
	}
	paging = *req.Pagination

	return
}

func (_i *galleryRepository) Create(gallery *schema.Gallery) (err error) {
	return _i.DB.DB.Create(&gallery).Error

}

func (_i *galleryRepository) Update(id uint64, gallery *schema.Gallery) (err error) {
	return _i.DB.DB.Model(&schema.Gallery{}).
		Where(&schema.Gallery{
			Model: gorm.Model{
				ID: uint(id),
			},
		}).
		Updates(gallery).Error
}

func (_i *galleryRepository) Delete(id uint64) (err error) {
	return _i.DB.DB.Delete(&schema.Gallery{}, id).Error
}
