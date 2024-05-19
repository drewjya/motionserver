package repository

import (
	"motionserver/app/database/schema"
	"motionserver/internal/bootstrap/database"
)

type bannerRepository struct {
	DB *database.Database
}

// Create implements BannerRepository.
func (_i *bannerRepository) Create(banner *schema.Banner) (err error) {
	old, err := _i.GetBanner()
	if err != nil {
		return err
	}
	if old != nil {
		return _i.DB.DB.Model(&old).Updates(banner).Error
	}
	return _i.DB.DB.Create(&banner).Error
}

// GetBanner implements BannerRepository.
func (_i *bannerRepository) GetBanner() (banner *schema.Banner, err error) {
	if err := _i.DB.DB.First(&banner).Error; err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
		return nil, nil
	}
	return banner, nil
}

type BannerRepository interface {
	GetBanner() (banner *schema.Banner, err error)
	Create(banner *schema.Banner) (err error)
}

func NewBannerRepository(db *database.Database) BannerRepository {
	return &bannerRepository{
		DB: db,
	}
}
