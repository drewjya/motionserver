package schema

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Tanggal datatypes.Date
	Berita  []byte
	Image   string `gorm:"type:varchar(255);not null"`
}
