package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ...
type Logger *zap.Logger

// NewLogger ...
func NewLogger(path string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{path}
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "ts"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.MessageKey = "message"

	return config.Build()
}
