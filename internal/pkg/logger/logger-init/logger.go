package loggerinit

import (
	"errors"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func InitLogger() error {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"logfile.log", "stdout"}

	var logger *zap.Logger
	logger, err := cfg.Build()
	if err != nil {
		return err
	}

	log = logger.Sugar()

	return nil
}

func GetLogger() (*zap.SugaredLogger, error) {
	if log != nil {
		return log, nil
	}

	return nil, errors.New("logger is not initialized")
}
