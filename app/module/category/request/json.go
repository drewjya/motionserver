package request

import (
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"
)

type CategoriesRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

func (req *CategoryRequest) ToDomain() (res *schema.Category) {
	res.Name = req.Name
	return
}
