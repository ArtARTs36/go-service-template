package cars

import (
	"context"

	carsapi "github.com/artarts36/go-service-template/pkg/cars-grpc-api/v1"
)

func (srv *Service) Get(
	ctx context.Context,
	req *carsapi.GetCarRequest,
) (*carsapi.Car, error) {
	return &carsapi.Car{}, nil
}
