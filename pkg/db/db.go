package db

import (
	"github.com/rs/zerolog"
	"github.com/vyantik/auth-jwt-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *config.DatabaseConfig, logger *zerolog.Logger) *Db {
	db, err := gorm.Open(postgres.Open(config.Url), &gorm.Config{})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
	}

	logger.Info().Msg("connected to database")

	return &Db{DB: db}
}