package config

import (
	"context"
	"log/slog"
	"os"

	"github.com/cappuccinotm/slogx"
	"github.com/cappuccinotm/slogx/slogm"
	"github.com/getsentry/sentry-go"
	slogmulti "github.com/samber/slog-multi"
	slogsentry "github.com/samber/slog-sentry/v2"
)

func (c *Container) setupLogger(conf Log) {
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: c.mapLogLevel(conf),
	})

	handlers := []slog.Handler{
		slogx.NewChain(jsonHandler, slogm.RequestID(), slogm.StacktraceOnError()),
	}

	if c.initSentry(conf) {
		handlers = append(handlers, slogsentry.Option{
			Level:     slog.LevelWarn,
			AddSource: true,
			AttrFromContext: []func(ctx context.Context) []slog.Attr{
				func(_ context.Context) []slog.Attr {
					return []slog.Attr{
						slog.String("release", c.appVersion),
					}
				},
				func(ctx context.Context) []slog.Attr {
					reqID, ok := slogm.RequestIDFromContext(ctx)
					if ok {
						return []slog.Attr{slog.Group(
							"tags",
							slog.String("request_id", reqID),
						)}
					}

					return []slog.Attr{}
				},
			},
		}.NewSentryHandler())
	}

	slog.SetDefault(slog.New(slogmulti.Fanout(handlers...)))
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

func (c *Container) initSentry(conf Log) bool {
	sentryInitErr := sentry.Init(sentry.ClientOptions{
		Dsn:           conf.Sentry.DSN,
		EnableTracing: false,
	})
	if sentryInitErr != nil {
		slog.
			With(slog.String("err", sentryInitErr.Error())).
			Error("failed to init sentry")
	}
	defer sentry.Flush(conf.Sentry.FlushTimeout)

	return sentryInitErr == nil
}
