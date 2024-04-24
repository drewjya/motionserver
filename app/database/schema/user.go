package schema

import (
	"database/sql/driver"
	"log"
	"motionserver/utils/helpers"
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Admin      Role = "admin"
	Superadmin Role = "superadmin"
	Basic      Role = "user"
)

func (role *Role) Scan(value interface{}) error {
	log.Println("scan role", value)
	if value == nil {
		return nil
	}
	if c, ok := value.([]byte); ok {

		*role = Role(c)
	} else if c, ok := value.(string); ok {
		*role = Role(c)
	}
	return nil
}

func (role Role) Value() (driver.Value, error) {
	return string(role), nil
}

type User struct {
	gorm.Model
	Account        Account   `gorm:"foreignKey:user_id"`
	Password       string    `gorm:"column:password;not null" json:"-"`
	Email          string    `gorm:"column:email;unique;not null" json:"email"`
	Role           Role      `sql:"type:role" gorm:"default:user"`
	LastAccessedAt time.Time `gorm:"column:last_accessed_at" json:"-"`
}

// compare password
func (u *User) ComparePassword(password string) bool {
	return helpers.CheckPasswordHash(password, u.Password)
}
