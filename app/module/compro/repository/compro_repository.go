package repository

import (
	"motionserver/app/database/schema"
	"motionserver/internal/bootstrap/database"
)

type comproRepository struct {
	DB *database.Database
}

// Create implements ComproRepository.
func (_i *comproRepository) Create(compro *schema.Compro) (err error) {
	old, err := _i.GetCompro()
	if err != nil {
		return err
	}
	if old != nil {
		return _i.DB.DB.Model(&old).Updates(compro).Error
	}
	return _i.DB.DB.Create(&compro).Error
}

// GetCompro implements ComproRepository.
func (_i *comproRepository) GetCompro() (compro *schema.Compro, err error) {
	if err := _i.DB.DB.First(&compro).Error; err != nil {
		return nil, err
	}
	return compro, nil
}

type ComproRepository interface {
	GetCompro() (compro *schema.Compro, err error)
	Create(compro *schema.Compro) (err error)
}

func NewComproRepository(db *database.Database) ComproRepository {
	return &comproRepository{
		DB: db,
	}
}
