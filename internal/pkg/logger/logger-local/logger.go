package loggerlocal

import (
	"fmt"

	"go.uber.org/zap"
)

func GetLogger() (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"logfile.log", "stdout"}

	var logger *zap.Logger
	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("logger is not initialized: %w", err)
	}

	return logger.Sugar(), nil
}
