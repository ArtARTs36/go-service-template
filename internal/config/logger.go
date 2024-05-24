package config

import (
	"log/slog"
	"os"

	"github.com/cappuccinotm/slogx"
	"github.com/cappuccinotm/slogx/slogm"
)

func (c *Container) setupLogger(conf Log) {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: c.mapLogLevel(conf),
	})

	logger := slog.New(
		slogx.Accumulator(
			slogx.NewChain(
				handler,
				slogm.RequestID(),
				slogm.StacktraceOnError(),
			),
		),
	)

	slog.SetDefault(logger)
}

func (c *Container) mapLogLevel(conf Log) slog.Level {
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

	return level
}
