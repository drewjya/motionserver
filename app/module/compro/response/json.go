package response

import (
	"motionserver/app/database/schema"
	"time"
)

// [Your response handling logic here]

type Compro struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Tanggal  string `json:"tanggal"`
	Data     string `json:"data"`
	Image    string `json:"image"`
	FileName string `json:"file_name"`
}

func FromDomain(domain *schema.Compro, image string) (compro *Compro) {
	if domain == nil {
		return nil
	}
	val, _ := domain.Tanggal.Value()
	format := val.(time.Time).Format("2006-01-02")
	return &Compro{
		ID:       domain.ID,
		Title:    domain.Title,
		Tanggal:  format,
		Subtitle: domain.Subtitle,
		FileName: domain.Image,
		Data:     string(domain.Data),
		Image:    image,
	}
}
