package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/vyantik/auth-jwt-service/config"
	"github.com/vyantik/auth-jwt-service/internal/auth"
	"github.com/vyantik/auth-jwt-service/internal/jwt"
	"github.com/vyantik/auth-jwt-service/internal/middleware"
	"github.com/vyantik/auth-jwt-service/internal/user"
	"github.com/vyantik/auth-jwt-service/pkg/db"
	"github.com/vyantik/auth-jwt-service/pkg/logger"
	"github.com/vyantik/auth-jwt-service/pkg/validators"
)

func main() {
	config.Init()
	customValidator := validators.NewValidator()

	//Configs
	//===============================================
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	serverConfig := config.NewServerConfig()
	dbConfig := config.NewDatabaseConfig()
	jwtConfig := config.NewJWTConfig()
	redisConfig := config.NewRedisConfig()
	//===============================================

	// Database
	//===============================================
	rdb := db.NewRedisDB(redisConfig, customLogger)
	db := db.NewDb(dbConfig, customLogger)
	//===============================================

	app := fiber.New(fiber.Config{
		AppName: "Jwt auth service",
		Prefork: false,
	})

	//Repositories
	//===============================================
	userRepository := user.NewRepository(db)
	//===============================================

	//Services
	//===============================================
	userService := user.NewService(userRepository)
	jwtService := jwt.NewService(jwtConfig.AccessSecret, jwtConfig.RefreshSecret)
	authService := auth.NewService(userService, jwtService)
	//===============================================

	//Middlewares
	//===============================================
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	//===============================================

	router := app.Group("/api")

	// Public routes (no authentication required)
	//===============================================
	auth.NewHandler(auth.HandlerDeps{
		Router:    router,
		Service:   authService,
		Logger:    customLogger,
		Validator: customValidator,
		Rdb:       rdb,
	})
	//===============================================

	// Protected routes (authentication required)
	//===============================================
	protected := router.Group("/profile", authMiddleware.Protected())
	protected.Get("/", func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		email := c.Locals("email").(string)

		return c.JSON(fiber.Map{
			"user_id": userID,
			"email":   email,
		})
	})
	//===============================================

	app.Listen(":" + serverConfig.Port)
}