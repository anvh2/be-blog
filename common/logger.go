package common

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// WrappedLogger ...
type WrappedLogger struct {
	logger *zap.Logger
}

// NewWrappedLogger ...
func NewWrappedLogger(logger *zap.Logger) *WrappedLogger {
	return &WrappedLogger{
		logger: logger,
	}
}

// Info ..
func (wlog *WrappedLogger) Info(msg string, fields ...zap.Field) {
	wlog.logger.Info(msg, fields...)
	fmt.Println(msg, fields)
}

// Error ..
func (wlog *WrappedLogger) Error(msg string, fields ...zap.Field) {
	wlog.logger.Error(msg, fields...)
	fmt.Println(msg, fields)
}

// Fatal ..
func (wlog *WrappedLogger) Fatal(msg string, fields ...zap.Field) {
	wlog.logger.Fatal(msg, fields...)
	fmt.Println(msg, fields)
}

// NewProductionZapLogger ...
func NewProductionZapLogger(path string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{path}
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "ts"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.MessageKey = "message"

	return config.Build()
}

// NewDevelopmentZapLogger ...
func NewDevelopmentZapLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
