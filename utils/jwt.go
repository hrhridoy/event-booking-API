package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string

func getEnvKey() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	secretKey = os.Getenv("secretHashKey")
}

func hashedSecretKey() (string, error) {
	getEnvKey()
	bytes, err := bcrypt.GenerateFromPassword([]byte(secretKey), 14)
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

func VerifyTOken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		errMsg := err.Error()
		return 0, errors.New(errMsg)
	}
	validToken := parsedToken.Valid
	if !validToken {
		return 0, errors.New("invalid token")
	}

	// Not required at the moment.
	// userId := getUserId(parsedToken.Claims.(jwt.MapClaims))

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// email := claims["email"].(string)
	userID := int64(claims["userId"].(float64))
	return userID, nil
}

// func getUserId(claims jwt.MapClaims) int64 {
// 	claim := claims
// 	// if !ok {
// 	// 	return 0, errors.New("invalid token claims")
// 	// }
// 	userId := claim["userId"].(int64)
// 	return userId
// }
