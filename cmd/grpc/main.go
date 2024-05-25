package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/artarts36/go-service-template/internal/port/grpc/app"

	log "github.com/sirupsen/logrus"
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
		log.Fatalln(err)
	}

	go func() {
		time.Sleep(1 * time.Minute)
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	slog.Info("[main] gracefully stopping")

	application.Stop()

	slog.Info("[main] gracefully stopped")
}
