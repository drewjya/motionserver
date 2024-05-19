package schema

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	Image    string
	Products []Product `gorm:"many2many:product_categories;"`
}
