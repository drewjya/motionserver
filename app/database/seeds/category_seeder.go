package seeds

import (
	"fmt"
	"motionserver/app/database/schema"
	"motionserver/internal/bootstrap/seeder"

	"gorm.io/gorm"
)

type CategorySeeder struct{}

func NewCategorySeeder() seeder.Seeder {
	return CategorySeeder{}
}

func (CategorySeeder) Seed(conn *gorm.DB) error {

	categories := []schema.Category{
		{
			Name: "T-Shirt",
			
		},
		{
			Name: "Shirt",
		},
	}
	for _, row := range categories {
		fmt.Println(row)
		if err := conn.Create(&row).Error; err != nil {
			return err
		}

	}
	return nil

}

func (CategorySeeder) Count(conn *gorm.DB) (int, error) {
	var count int64
	if err := conn.Model(&schema.Category{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
