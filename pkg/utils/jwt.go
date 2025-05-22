package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var accessSecret = []byte("access-secret-key")
// var refreshSecret = []byte("refresh-secret-key")

func GenerateToken(userId uint, duration time.Duration, secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func VerifyToken(tokenStr string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
