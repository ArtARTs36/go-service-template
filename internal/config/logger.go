package config

import (
	"log/slog"
	"os"
)

func (c *Container) setupLogger(conf *LogConfig) {
	var level slog.Level

	switch conf.Level {
	case "info":
		level = slog.LevelInfo
	case "warning":
	case "warn":
		level = slog.LevelWarn
	case "error":
	case "err":
		level = slog.LevelError
	default:
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	slog.SetDefault(logger)
}
