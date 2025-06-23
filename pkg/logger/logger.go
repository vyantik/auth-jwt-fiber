package logger

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/vyantik/auth-jwt-service/config"
)

func NewLogger(config *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(config.Level)
	var logger zerolog.Logger
	
	if config.Format == "json" {
		// Создаем папку logs если её нет
		logsDir := "logs"
		if err := os.MkdirAll(logsDir, 0755); err != nil {
			// Если не удалось создать папку, используем stderr
			logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		} else {
			// Создаем файл для логов
			logFile, err := os.OpenFile(
				filepath.Join(logsDir, "app.log"),
				os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				0644,
			)
			if err != nil {
				// Если не удалось открыть файл, используем stderr
				logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
			} else {
				logger = zerolog.New(logFile).With().Timestamp().Logger()
			}
		}
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}
	return &logger
}