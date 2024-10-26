package utils

import (
	bcrypt "golang.org/x/crypto/bcrypt"
)

func HashSaltedPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifySaltedPassword(pass string, passHash string) (bool, error) {
	ok := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(pass))
	if ok != nil {
		return false, ok
	}
	return true, nil
}
