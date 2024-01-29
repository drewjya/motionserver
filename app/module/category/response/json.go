package response

import "motionserver/app/database/schema"

type Categories struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FromDomain(category *schema.Category) (res *Categories) {
	if category != nil {
		res = &Categories{
			ID:   category.ID,
			Name: category.Name,
		}
	}
	return
}
