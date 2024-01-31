package request

import (
	"mime/multipart"
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"
	"time"

	"gorm.io/datatypes"
)

// [Your request handling logic here]

type GalleriesRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

type GalleryRequest struct {
	Title   string    `form:"title" json:"title" validate:"required"`
	Tanggal time.Time `form:"tanggal" json:"tanggal" validate:"required"`
	Detail  string    `form:"detail" json:"detail" validate:"required"`
	Image   string
	File    *multipart.FileHeader
}

func (req *GalleryRequest) ToDomain() (r *schema.Gallery) {
	res := new(schema.Gallery)
	date := datatypes.Date(req.Tanggal)
	res.Title = req.Title
	res.Tanggal = date
	res.Detail = req.Detail
	res.Image = req.Image
	return res
}
