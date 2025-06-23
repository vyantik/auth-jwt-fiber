package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	log.Info("Loaded .env file")
}

func NewDatabaseConfig() *DatabaseConfig {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", GetString("POSTGRES_USER", ""), GetString("POSTGRES_PASSWORD", ""), GetString("POSTGRES_HOST", ""), GetString("POSTGRES_PORT", ""), GetString("POSTGRES_DB", ""))
	return &DatabaseConfig{
		Url: url,
	}
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: GetString("APPLICATION_PORT", "3000"),
	}
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  zerolog.Level(GetInt("LOG_LEVEL", 0)),
		Format: GetString("LOG_FORMAT", "json"),
	}
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		AccessSecret:  GetString("JWT_ACCESS_SECRET", ""),
		RefreshSecret: GetString("JWT_REFRESH_SECRET", ""),
	}
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     GetString("REDIS_HOST", "localhost"),
		Port:     GetInt("REDIS_PORT", 6379),
		Password: GetString("REDIS_PASSWORD", ""),
		DB:       GetInt("REDIS_DB", 0),
	}
}

func GetInt(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

func GetBool(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

func GetString(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
