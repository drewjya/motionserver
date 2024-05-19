package schema

import "gorm.io/gorm"

type PromotionProduct struct {
	gorm.Model
	ProductID uint    // Foreign key
	Product   Product `gorm:"constraint:OnDelete:SET NULL;"`
	Image     string
}
