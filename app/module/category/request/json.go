package request

import (
	"log"
	"motionserver/app/database/schema"
	"motionserver/utils/paginator"
)

type CategoriesRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

type CategoryRequest struct {
	Name string `form:"name" json:"name" validate:"required,min=3,max=255"`
}

func (req *CategoryRequest) ToDomain() (res *schema.Category) {
	log.Println("req", req)
	return &schema.Category{
		Name: req.Name,
	}
}
