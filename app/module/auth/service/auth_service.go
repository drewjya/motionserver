package service

import (
	"errors"
	"motionserver/app/database/schema"
	"motionserver/app/middleware"
	"motionserver/app/module/auth/repository"
	"motionserver/app/module/auth/request"
	"motionserver/app/module/auth/response"
	"motionserver/utils/helpers"
	"time"
)

type authService struct {
	Repo repository.AuthRepository
}

type AuthService interface {
	Login(req request.LoginRequest) (res response.LoginResponse, err error)
	Register(req request.RegisterRequest) (res response.RegisterResponse, err error)
	Profile(userId uint) (res response.LoginResponse, err error)
	RefreshToken(userId uint64) (res *middleware.TokenResponse, err error)
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		Repo: repo,
	}
}

func (_i *authService) RefreshToken(userId uint64) (res *middleware.TokenResponse, err error) {
	res, err = middleware.GenerateTokenUser(middleware.TokenData{
		UserId: userId,
		Roles:  "user",
	})
	return
}

func (_i *authService) Profile(userId uint) (res response.LoginResponse, err error) {
	user, err := _i.Repo.FindUserByUserId(userId)
	if err != nil {
		return
	}
	res.Name = user.Account.Name
	res.Email = user.Email
	res.UserId = uint64(user.ID)
	res.Role = string(user.Role)
	res.AccountId = uint64(user.Account.ID)
	return
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
		Roles:  string(user.Role),
	})

	if err != nil {
		return
	}

	res.Name = user.Account.Name
	res.Email = user.Email
	res.Role = string(user.Role)
	res.UserId = uint64(user.ID)
	res.AccountId = uint64(user.Account.ID)

	res.Token = *resp

	return

}

func (_i *authService) Register(req request.RegisterRequest) (res response.RegisterResponse, err error) {
	// log.Println(*req.Role, "role")
	if req.Role == nil {
		user := schema.Basic
		req.Role = &user
	}

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
		Role:           *req.Role,
		Account:        schema.Account{Name: req.Name},
		LastAccessedAt: time.Now(),
	}

	user, err = _i.Repo.CreateUser(&newUser)

	res.Name = user.Account.Name
	res.Email = user.Email
	res.UserId = uint64(user.ID)
	res.AccountId = uint64(user.Account.ID)
	res.Role = string(user.Role)
	return

}
