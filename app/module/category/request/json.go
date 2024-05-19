package request

import (
	"mime/multipart"
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"
)

type CategoriesRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

type CategoryRequest struct {
	Name  string `form:"name" json:"name" validate:"required,min=3,max=255"`
	Image string
	File  *multipart.FileHeader
}



func (req *CategoryRequest) ToDomain() (res *schema.Category) {
	if req == nil {
		return nil
	}
	return &schema.Category{
		Name:  req.Name,
		Image: req.Image,
	}
}
