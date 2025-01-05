package bootstrap

import (
	"go.uber.org/zap"
)

const logFilePath = "logfile"

func mustInitLogger() {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout", logFilePath},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	if err := logger.Sync(); err != nil {
		zap.L().Error("Failed to sync logger", zap.Error(err))
	}

	zap.ReplaceGlobals(logger)

	zap.L().Info("Custom logger initialized")
}
