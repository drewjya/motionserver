package schema

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Compro struct {
	gorm.Model
	Title    string
	Subtitle string
	Tanggal  datatypes.Date
	Image    string
	Data []byte
}
