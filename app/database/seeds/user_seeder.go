package seeds

import (
	"motionserver/app/database/schema"
	"motionserver/internal/bootstrap/seeder"

	"motionserver/utils/helpers"
	"time"

	"gorm.io/gorm"
)

type UserSeeder struct{}

func NewUserSeeder() seeder.Seeder {
	return UserSeeder{}
}

func (UserSeeder) Seed(conn *gorm.DB) error {
	password, err := helpers.Hash("password")
	if err != nil {
		panic(err)
	}
	var users = []schema.User{
		{
			Email:          "andre@email.com",
			Password:       password,
			LastAccessedAt: time.Now(),
			Account: schema.Account{
				Name: "Andre",
			},
			Role: schema.Superadmin,
		},
		{
			Email:          "william@email.com",
			Password:       password,
			LastAccessedAt: time.Now(),
			Account: schema.Account{
				Name: "William",
			},
		},
	}

	for _, row := range users {
		if err = conn.Create(&row).Error; err != nil {
			return err
		}

	}
	return nil

}

func (UserSeeder) Count(conn *gorm.DB) (int, error) {
	var count int64
	if err := conn.Model(&schema.User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
