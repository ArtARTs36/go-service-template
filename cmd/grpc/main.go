package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/artarts36/go-service-template/internal/port/grpc/app"
)

var (
	Version = "0.1.0" //nolint: gochecknoglobals // version is need
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	slog.Info("[main] starting")

	cfg, err := app.InitConfig("CARS_")
	if err != nil {
		slog.
			With(slog.Any("err", err)).
			Error("[main] failed to initialize config")

		os.Exit(1)
	}

	application, err := app.NewApp(cfg, Version)
	if err != nil {
		slog.
			With(slog.Any("err", err)).
			Error("[main] failed to initialize application")

		os.Exit(1)
	}

	go func() {
		if appRunErr := application.Run(); appRunErr != nil {
			slog.
				With(slog.Any("err", appRunErr.Error())).
				Error("[main] failed to run application")
		}
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	slog.Info("[main] gracefully stopping")

	application.Stop()

	slog.Info("[main] gracefully stopped")
}
