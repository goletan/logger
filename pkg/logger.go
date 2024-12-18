package pkg

import (
	"github.com/goletan/logger/internal/logger"
	"go.uber.org/zap"
	"log"
)

func NewLogger() (*logger.ZapLogger, error) {
	newLogger, err := logger.NewLogger()
	if err != nil {
		log.Fatal("Failed to create logger", zap.Error(err))
		return nil, err
	}

	return newLogger, nil
}
