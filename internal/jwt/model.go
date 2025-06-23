package jwt

import "github.com/golang-jwt/jwt/v5"

type Service struct {
	accessSecret  string
	refreshSecret string
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
