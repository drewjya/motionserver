package response

import "motionserver/app/database/schema"

type Categories struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func FromDomain(category *schema.Category, image string) (res *Categories) {
	if category != nil {
		res = &Categories{
			ID:    category.ID,
			Image: image,

			Name: category.Name,
		}
	}
	return
}

func FromDomainNo(category *schema.Category) (res *Categories) {
	if category != nil {
		res = &Categories{
			ID:    category.ID,

			Name: category.Name,
		}
	}
	return
}
