package model

import (
	"errors"

	"github.com/GareArc/MovieMate/internal/utils"
)

type User struct {
	BaseModel
	Email    string `gorm:"column:email;length:255;not null;unique"`
	Password string `gorm:"column:password;length:255;not null"`
	Nickname string `gorm:"column:nickname;length:255"`
	Avatar   string `gorm:"column:avatar;length:255"`
}

func (u *User) VerifyPassword(givenPass string) (bool, error) {
	hashedPass := u.Password
	res, ok := utils.VerifySaltedPassword(givenPass, hashedPass)
	if ok != nil {
		return false, errors.New("password not match")
	}
	return res, nil
}

func (u *User) SetPassword(pass string) error {
	hashedPass, err := utils.HashSaltedPassword(pass)
	if err != nil {
		return err
	}
	u.Password = hashedPass
	return nil
}
