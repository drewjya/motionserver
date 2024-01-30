package response

import (
	"motionserver/app/database/schema"
	"time"
)

// [Your response handling logic here]

type Gallery struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Tanggal string `json:"tanggal"`
	Detail  string `json:"detail"`
	Image   string `json:"image"`
}

func FromDomain(domain *schema.Gallery) (gallery *Gallery) {
	if domain == nil {
		return nil
	}
	val, _ := domain.Tanggal.Value()
	format := val.(time.Time).Format("2006-01-02")
	return &Gallery{
		ID:      domain.ID,
		Title:   domain.Title,
		Tanggal: format,
		Detail:  domain.Detail,
		Image:   domain.Image,
	}
}
