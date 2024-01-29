package schema

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name   string `gorm:"column:name;not null"`
	UserId uint   `gorm:"column:user_id;not null"`
}
