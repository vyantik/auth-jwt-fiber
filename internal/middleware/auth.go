package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vyantik/auth-jwt-service/internal/jwt"
)

func NewAuthMiddleware(jwtService *jwt.Service) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "missing authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid authorization header format")
		}

		token := parts[1]
		claims, err := m.jwtService.ValidateAccessToken(token)
		if err != nil {
			if err == jwt.ErrExpiredToken {
				return fiber.NewError(fiber.StatusUnauthorized, "token has expired")
			}
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
		}

		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}