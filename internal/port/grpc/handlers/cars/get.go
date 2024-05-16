package cars

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/artarts36/go-service-template/internal/application/car"
	carsapi "github.com/artarts36/go-service-template/pkg/cars-grpc-api"
)

func (s *Service) Get(ctx context.Context, req *carsapi.GetCarRequest) (*carsapi.Car, error) {
	if err := s.validateGetCarRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	c, err := s.container.Application.Operations.Car.Get.Get(ctx, &car.GetOperationParams{
		ID: req.Id,
	})
	if err != nil {
		if errors.Is(err, car.ErrCarNotFound) {
			return nil, status.Error(codes.NotFound, "car not found")
		}

		return nil, status.Error(codes.Internal, "unable to get car")
	}

	return &carsapi.Car{
		Id: c.ID,
	}, nil
}

func (s *Service) validateGetCarRequest(req *carsapi.GetCarRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required),
	}
}
