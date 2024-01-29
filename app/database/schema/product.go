package schema

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       uint64
	Stock       uint
	Image       string
	SerialCode  string
	Categories  []Category `gorm:"many2many:product_categories;"`
}
