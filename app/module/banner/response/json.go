package response

import (
	"motionserver/app/database/schema"
)

// [Your response handling logic here]

type Banner struct {
	Label     string `json:"string"`
	Image     string `json:"image"`
	ImageName string `json:"image_name"`
}

func FromDomain(domain *schema.Banner, image string) (compro *Banner) {
	if domain == nil {
		return nil
	}

	return &Banner{
		Label:     domain.Label,
		Image:     image,
		ImageName: domain.Image,
	}
}
