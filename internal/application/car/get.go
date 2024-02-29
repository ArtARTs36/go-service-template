package car

import (
	"context"

	"go-service-template/internal/domain"
)

type GetOperation struct {
	cars domain.CarRepository
}

type GetOperationParams struct {
	ID int
}

func NewGetOperation(cars domain.CarRepository) *GetOperation {
	return &GetOperation{cars: cars}
}

func (o *GetOperation) Get(ctx context.Context, params *GetOperationParams) (*domain.Car, error) {
	return o.cars.Find(ctx, params.ID)
}
