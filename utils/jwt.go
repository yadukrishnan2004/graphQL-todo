package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("supersecret") 

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
}