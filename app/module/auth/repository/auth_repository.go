package repository

import (
	"log"
	"template_clean_architecture/app/database/schema"
	"template_clean_architecture/internal/bootstrap/database"
	"time"
)

type authRepository struct {
	DB *database.Database
}

type AuthRepository interface {
	FindUserByEmail(email string) (user *schema.User, err error)
	CreateUser(user *schema.User) (res *schema.User, err error)
	UpdateLastLogin(user *schema.User) (res *schema.User, err error)

	FindAccountByUserId(userId uint) (account *schema.Account, err error)
}

func NewAuthRepository(db *database.Database) AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (_i *authRepository) FindAccountByUserId(userId uint) (account *schema.Account, err error) {
	if err := _i.DB.DB.Where("user_id = ?", userId).First(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (_i *authRepository) FindUserByEmail(email string) (user *schema.User, err error) {
	if err := _i.DB.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println(err, "asjsja")
		return nil, err
	}

	return user, nil

}
func (_i *authRepository) CreateUser(user *schema.User) (res *schema.User, err error) {
	if err := _i.DB.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *authRepository) UpdateLastLogin(user *schema.User) (res *schema.User, err error) {
	user.LastAccessedAt = time.Now()

	if err := _i.DB.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}
