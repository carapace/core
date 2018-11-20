package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Logger implements a global logger which may be imported by packages implementing components for the core suite.
//
// Dependencies with special logging needs (for example because they output a very large quantity of log statements)
// should explicitly request a logger object in their constructors. If a dependency does not require a logging object;
// it should implement core.Logger or no logging at all.
//
// An error should only be logged if the function handles it. If it is simply returned; it is up to the caller to log the
// error and it's context.
var Logger *zap.Logger

func init() {
	Logger = defaultLogger("INFO")
}

const (
	// LogKey is the default key used to distinguish loggers.
	LogKey = "logger"
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
	logger = logger.With(zap.String(LogKey, "core.Logger"))
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
