package logger

import (
	"context"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerContextKeyType string

const loggerContextKey loggerContextKeyType = "logger"

var Logger *zap.Logger

// initializes logger
func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.Sampling = nil
	cfg.OutputPaths = []string{"stdout"}

	var err error
	if Logger, err = cfg.Build(); err != nil {
		panic(err)
	}

	defer func() {
		if err := Logger.Sync(); err != nil {
			log.Println("Failed to sync logger: " + err.Error())
		}
	}()
}

// Inject return context including the given logger
func Inject(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// Ctx return Logger with context including request ID
func Ctx(ctx context.Context) *zap.Logger {
	if logger := ctx.Value(loggerContextKey); logger != nil {
		return logger.(*zap.Logger)
	}
	return Logger
}
