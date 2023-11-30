package middleware

import (
	"fmt"

	"go.uber.org/zap"
)

func InitLogger(devMode bool) error {
	var (
		logger *zap.Logger
		err    error
	)
	if devMode {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return fmt.Errorf("failed to start development logger, error: %v", err)
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			return fmt.Errorf("failed to start production logger, error: %v", err)
		}
	}

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
	return nil
}
