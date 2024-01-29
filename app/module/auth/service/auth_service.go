package service

import (
	"errors"
	"template_clean_architecture/app/database/schema"
	"template_clean_architecture/app/middleware"
	"template_clean_architecture/app/module/auth/repository"
	"template_clean_architecture/app/module/auth/request"
	"template_clean_architecture/app/module/auth/response"
	"template_clean_architecture/utils/helpers"
	"time"
)

type authService struct {
	Repo repository.AuthRepository
}

type AuthService interface {
	Login(req request.LoginRequest) (res response.LoginResponse, err error)
	Register(req request.RegisterRequest) (res response.RegisterResponse, err error)
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		Repo: repo,
	}
}

func (_i *authService) Login(req request.LoginRequest) (res response.LoginResponse, err error) {
	user, err := _i.Repo.FindUserByEmail(req.Email)
	if err != nil {
		return
	}
	if user == nil {
		return
	}
	if !user.ComparePassword(req.Password) {
		err = errors.New("password not match")
		return
	}
	user, err = _i.Repo.UpdateLastLogin(user)
	if err != nil {
		return
	}

	account, err := _i.Repo.FindAccountByUserId(user.ID)
	if err != nil {
		return
	}
	user.Account = *account

	resp, err := middleware.GenerateTokenUser(middleware.TokenData{
		UserId: uint64(user.ID),
		Roles:  "user",
	})

	if err != nil {
		return
	}

	res.Name = user.Account.Name
	res.Email = user.Email
	res.UserId = uint64(user.ID)
	res.AccountId = uint64(user.Account.ID)

	res.Token = *resp

	return

}

func (_i *authService) Register(req request.RegisterRequest) (res response.RegisterResponse, err error) {
	user, err := _i.Repo.FindUserByEmail(req.Email)
	if err != nil && err.Error() != "record not found" {
		return
	}

	if user != nil {
		err = errors.New("email already exists")
		return
	}

	newPassword, err := helpers.Hash(req.Password)
	if err != nil {
		return
	}
	newUser := schema.User{
		Email:          req.Email,
		Password:       newPassword,
		Account:        schema.Account{Name: req.Name},
		LastAccessedAt: time.Now(),
	}

	user, err = _i.Repo.CreateUser(&newUser)

	res.Name = user.Account.Name
	res.Email = user.Email
	res.UserId = uint64(user.ID)
	res.AccountId = uint64(user.Account.ID)

	return

}
