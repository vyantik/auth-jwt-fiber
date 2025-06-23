package config

import "github.com/rs/zerolog"

type DatabaseConfig struct {
	Url string
}

type ServerConfig struct {
	Port string
}

type LogConfig struct {
	Level  zerolog.Level
	Format string
}

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}
