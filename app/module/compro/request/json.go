package request

import (
	"mime/multipart"
	"motionserver/app/database/schema"
	"time"

	"gorm.io/datatypes"
)

// [Your request handling logic here]

type ComproRequest struct {
	Title    string    `form:"title" json:"title" validate:"required"`
	Subtitle string    `form:"subtitle" json:"subtitle" validate:"required"`
	Tanggal  time.Time `form:"tanggal" json:"tanggal" validate:"required"`
	Data     string    `form:"data" json:"data" validate:"required"`
	Image    string
	File     *multipart.FileHeader
}

func (req *ComproRequest) ToDomain() (r *schema.Compro) {
	res := new(schema.Compro)
	date := datatypes.Date(req.Tanggal)
	res.Title = req.Title
	res.Tanggal = date
	res.Subtitle = req.Subtitle
	res.Data = []byte(req.Data)
	res.Image = req.Image
	return res
}
