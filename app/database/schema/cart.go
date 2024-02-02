package schema

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	AccountID uint    `gorm:"column:account_id"`
	Quantity  int32   `gorm:"column:quantity"`
	ProductID uint    `gorm:"column:product_id"`
	Product   Product `gorm:"foreignKey:product_id"`
}
