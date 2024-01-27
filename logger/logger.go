package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func init() {
	logLevel := new(slog.LevelVar)
	logLevel.Set(slog.LevelDebug)
	ops := slog.HandlerOptions{
		Level: logLevel,
	}

	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &ops))
	slog.SetDefault(Logger)
	Logger.Info("Init logger.")
}
