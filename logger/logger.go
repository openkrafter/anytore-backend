package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

type LogConfig struct {
	LOG_LEVEL string `yaml:"LOG_LEVEL"`
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func InitLogger() {
	logConfig := new(LogConfig)

	// env: LOG_LEVEL, value: debug or info, default: debug
	logConfig.LOG_LEVEL = "debug"
	if os.Getenv("LOG_LEVEL") == "debug" || os.Getenv("LOG_LEVEL") == "info" {
		logConfig.LOG_LEVEL = os.Getenv("LOG_LEVEL")
	}

	logLevel := new(slog.LevelVar)
	if logConfig.LOG_LEVEL == "info" {
		logLevel.Set(slog.LevelInfo)
	} else if logConfig.LOG_LEVEL == "debug" {
		logLevel.Set(slog.LevelDebug)
	}

	ops := slog.HandlerOptions{
		Level: logLevel,
	}

	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &ops))
	slog.SetDefault(Logger)
	Logger.Debug("Init logger.")

}
