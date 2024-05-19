package schema

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Image string
	Label string
}
