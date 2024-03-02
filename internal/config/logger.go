package config

import (
	"log/slog"
	"os"
)

func (c *Container) setupLogger(conf *LogConfig) {
	var level slog.Level

	if conf.Level == "info" {
		level = slog.LevelInfo
	} else if conf.Level == "warn" || conf.Level == "warning" {
		level = slog.LevelWarn
	} else if conf.Level == "error" || conf.Level == "err" {
		level = slog.LevelError
	} else {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	slog.SetDefault(logger)
}
