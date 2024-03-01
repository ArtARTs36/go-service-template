package main

import (
	"github.com/artarts36/go-service-template/internal/port/grpc"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := grpc.InitConfig("CARS_")
	if err != nil {
		log.Fatalln(err)
	}

	application := grpc.NewApp(cfg)

	go func() {
		if err := application.Run(); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	log.Info("Gracefully stopped")
}
