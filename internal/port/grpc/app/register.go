package app

import (
	"google.golang.org/grpc"

	"github.com/artarts36/go-service-template/internal/config"
	"github.com/artarts36/go-service-template/internal/port/grpc/handlers/cars"
	carsapi "github.com/artarts36/go-service-template/pkg/cars-grpc-api"
)

func registerServices(gRPCServer grpc.ServiceRegistrar, cont *config.Container) {
	carsapi.RegisterCarsServiceServer(gRPCServer, cars.NewService(cont))
}
