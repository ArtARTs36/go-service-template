package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/artarts36/go-service-template/internal/port/grpc/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := app.InitConfig("CARS_")
	if err != nil {
		log.Fatalln(err)
	}

	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if runErr := application.Run(); runErr != nil {
			panic(runErr)
		}
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	log.Info("Gracefully stopped")
}
