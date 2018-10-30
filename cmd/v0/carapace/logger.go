package carapace

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

func defaultLogger(level string) *zap.Logger {
	lvl := parseLogLevel(level)

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(lvl),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		// This panic should only occur during development, when reconfiguring the logger, or when
		// switching between major version of zap.
		panic(fmt.Sprintf("default logger incorrectly configured: %s", err.Error()))
	}
	logger = logger.With(zap.String("LOGGER", "CARAPACE-DEFAULT"))
	return logger
}

func parseLogLevel(level string) zapcore.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
