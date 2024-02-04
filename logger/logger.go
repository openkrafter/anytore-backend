package logger

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

var Logger *slog.Logger

type LogConfig struct {
	LOG_LEVEL string `yaml:"LOG_LEVEL"`
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func init() {
	logConfig := new(LogConfig)
	defaultConfig, err := os.ReadFile("config/defaultConfig.yaml")
	if err != nil {
		slog.Error("Failed to read defaultConfig.yaml.", ErrAttr(err))
	}
	yaml.Unmarshal(defaultConfig, &logConfig)

	if os.Getenv("LOG_LEVEL") != "" {
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
