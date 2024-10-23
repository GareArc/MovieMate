package utils

import (
	"errors"
	"time"

	"github.com/GareArc/MovieMate/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func JWTCreateToken(user_id string) (string, error) {
	config := config.GetStaticConfig()
	expire_hours := config.Int("jwt.expire")
	jwt_secret := []byte(config.String("jwt.secret"))

	if expire_hours == 0 || jwt_secret == nil {
		return "", errors.New("jwt secret or expire time not set")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user_id,
			"exp":     time.Now().Add(time.Duration(expire_hours) * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString(jwt_secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	config := config.GetStaticConfig()
	jwt_secret := []byte(config.String("jwt.secret"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwt_secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
