package response

import (
	"motionserver/app/database/schema"
	"time"
)

// [Your response handling logic here]

type News struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Tanggal string `json:"tanggal"`
	Berita  string `json:"detail"`
	Image   string `json:"image"`
}

func FromDomain(domain *schema.News, image string) (cart *News) {
	if domain == nil {
		return nil
	}
	val, _ := domain.Tanggal.Value()
	format := val.(time.Time).Format("2006-01-02")
	return &News{
		ID:      domain.ID,
		Title:   domain.Title,
		Tanggal: format,
		Berita:  string(domain.Berita),
		Image:   image,
	}
}
