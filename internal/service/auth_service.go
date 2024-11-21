package service

import (
	"errors"

	"github.com/GareArc/MovieMate/internal/db"
	"github.com/GareArc/MovieMate/internal/types/model"
	"github.com/GareArc/MovieMate/internal/utils"
)

type AuthService struct{}

func (as *AuthService) Login(email string, password string) (string, *model.User, error) {
	var user model.User
	db.MainDB.Model(&model.User{}).Where("email = ?", email).First(&user)

	if user.ID == "" {
		return "", nil, errors.New("user not found")
	}

	res, err := user.VerifyPassword(password)
	if err != nil {
		return "", nil, err
	}

	if !res {
		return "", nil, errors.New("password not match")
	}

	jwt_token, err := utils.JWTCreateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return jwt_token, &user, nil
}

func (as *AuthService) Register(email string, password string, nickname string) (string, *model.User, error) {
	var user model.User
	db.MainDB.Model(&model.User{}).Where("email = ?", email).First(&user)

	if user.ID != "" {
		return "", nil, errors.New("user with the same email already exists")
	}

	user = model.User{
		Email:    email,
		Nickname: nickname,
	}

	user.SetPassword(password)

	db.MainDB.Create(&user)

	jwt_token, err := utils.JWTCreateToken(user.ID)

	return jwt_token, &user, err
}
