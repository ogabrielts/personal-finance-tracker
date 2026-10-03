package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("KEY")

type Claims struct {
	UserId int `json:"user_id"`
	UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, name, email string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)

	claim := Claims {
		UserId: id,
		UserName: name,
		UserEmail: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(jwtKey))
}

func ValidateToken(tokenString string) (*Claims, error) {
	claim := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (any, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claim, nil
}
