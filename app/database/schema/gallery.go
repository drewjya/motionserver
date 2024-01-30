package schema

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Gallery struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Tanggal datatypes.Date
	Detail  string `gorm:"type:varchar(255);not null"`
	Image   string `gorm:"type:varchar(255);not null"`
}
