package request

import (
	"mime/multipart"
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"
	"time"

	"gorm.io/datatypes"
)

// [Your request handling logic here]

type NewssRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

type NewsRequest struct {
	Title   string    `form:"title" json:"title" validate:"required"`
	Tanggal time.Time `form:"tanggal" json:"tanggal" validate:"required"`
	Berita  string    `form:"berita" json:"berita" validate:"required"`
	Image   string
	File    *multipart.FileHeader
}

func (req *NewsRequest) ToDomain() (r *schema.News) {
	res := new(schema.News)
	date := datatypes.Date(req.Tanggal)

	res.Title = req.Title
	res.Tanggal = date
	res.Berita = []byte(req.Berita)
	res.Image = req.Image

	return res
}
