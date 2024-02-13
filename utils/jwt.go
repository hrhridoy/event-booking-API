package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string

func hashedSecretKey() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("Hr_Hr1doy"), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func GenerateToken(email string, userId int64) (string, error) {
	if secretKey == "" {
		hash, err := hashedSecretKey()
		if err != nil {
			return "", err
		}
		secretKey = hash
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
