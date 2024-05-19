package request

import (
	"mime/multipart"
	"motionserver/app/database/schema"
)

// [Your request handling logic here]

type BannerRequest struct {
	Label string `form:"label" json:"label" validate:"required"`
	Image string
	File  *multipart.FileHeader
}

func (req *BannerRequest) ToDomain() (r *schema.Banner) {
	res := new(schema.Banner)

	res.Label = req.Label
	res.Image = req.Image
	return res
}
