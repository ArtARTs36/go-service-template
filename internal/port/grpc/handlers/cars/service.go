package cars

import (
	"github.com/artarts36/go-service-template/internal/config"
	carsapi "github.com/artarts36/go-service-template/pkg/cars-grpc-api"
)

type Service struct {
	carsapi.UnimplementedCarsServiceServer

	container *config.Container
}

func NewService(cont *config.Container) *Service {
	return &Service{container: cont}
}
