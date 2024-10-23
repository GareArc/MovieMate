package utils

import (
	bcrypt "golang.org/x/crypto/bcrypt"
)

func HashSaltedPassword(pass string) (string, string, error) {
	salt := GenerateRandomString(16)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass+salt), bcrypt.DefaultCost)
	return string(hash), salt, err
}

func VerifySaltedPassword(pass string, salt string, passHash string) (bool, error) {
	hash, ok := bcrypt.GenerateFromPassword([]byte(pass+salt), bcrypt.DefaultCost)
	if ok != nil {
		return false, ok
	}
	return string(hash) == passHash, nil
}
