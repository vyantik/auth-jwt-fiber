package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/vyantik/auth-jwt-service/pkg/validators"
)

type HandlerDeps struct {
	Router    fiber.Router
	Service   *Service
	Logger    *zerolog.Logger
	Validator *validators.Validator
	Rdb       *redis.Client
}

type Handler struct {
	router    fiber.Router
	service   *Service
	logger    *zerolog.Logger
	validator *validators.Validator
	rdb       *redis.Client
}

func NewHandler(deps HandlerDeps) {
	h := Handler{
		router:    deps.Router,
		service:   deps.Service,
		logger:    deps.Logger,
		validator: deps.Validator,
		rdb:       deps.Rdb,
	}

	h.router.Post("/register", h.Register)
	h.router.Post("/login", h.Login)
	h.router.Post("/refresh", h.Refresh)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	req := RegisterRequest{
		Email:    c.FormValue("email"),
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	if err := h.validator.ValidateRequest(c, req); err != nil {
		return err
	}

	_, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to register",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request",
		})
	}

	if err := h.validator.ValidateRequest(c, req); err != nil {
		return err
	}

	tokenPair, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect email or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"data":    tokenPair,
	})
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request",
		})
	}

	if err := h.validator.ValidateRequest(c, req); err != nil {
		return err
	}

	tokenPair, err := h.service.RefreshTokens(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid refresh token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tokens refreshed successfully",
		"data":    tokenPair,
	})
}