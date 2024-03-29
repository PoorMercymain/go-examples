package loggeronce

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type logger struct {
	*zap.SugaredLogger
	once *sync.Once
}

var log = &logger{once: &sync.Once{}}

func Logger() *logger {
	var err error
	log.once.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.OutputPaths = []string{"logfile.log", "stdout"}

		var logger *zap.Logger
		logger, err = cfg.Build()

		log.SugaredLogger = logger.Sugar()
	})

	if err != nil {
		panicStr := fmt.Errorf("cannot proceed: couldn't initialize logger because of an error - %w", err)
		panic(panicStr)
	}

	return log
}
