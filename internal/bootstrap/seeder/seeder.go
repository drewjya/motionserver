package seeder

import "gorm.io/gorm"

type Seeder interface {
	Seed(*gorm.DB) error
	Count(*gorm.DB) (int, error)
}
