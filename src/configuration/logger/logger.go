package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func getOutputLog() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLogLevel() zapcore.Level {
	level := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL)))
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLog()},
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.EpochMillisTimeEncoder,
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build()
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}
func Warn(msg string, tags ...zap.Field) {
	log.Warn(msg, tags...)
	log.Sync()
}
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}
