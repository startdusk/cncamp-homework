package logger

import (
	"go.uber.org/zap"
)

func New(logFile, logLevel string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.OutputPaths = []string{logFile}

	switch logLevel {
	case "DEBUG":
		cfg.Level.SetLevel(zap.DebugLevel)
	case "WARNING":
		cfg.Level.SetLevel(zap.WarnLevel)
	case "ERROR":
		cfg.Level.SetLevel(zap.ErrorLevel)
	default:
		cfg.Level.SetLevel(zap.InfoLevel)
	}
	return cfg.Build()
}
