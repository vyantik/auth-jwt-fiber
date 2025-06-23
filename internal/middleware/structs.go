package middleware

import (
	"github.com/vyantik/auth-jwt-service/internal/jwt"
)

type AuthMiddleware struct {
	jwtService *jwt.Service
}