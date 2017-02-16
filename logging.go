package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger() (*zap.Logger, error) {
	logConfig := zap.NewProductionConfig()
	logConfig.Encoding = "console"
	logConfig.DisableCaller = true
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := logConfig.Build()
	return logger, err
}
