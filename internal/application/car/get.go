package car

import (
	"context"
	"errors"

	"github.com/artarts36/go-service-template/internal/domain"
	"github.com/artarts36/go-service-template/internal/infrastructure/repository"
)

type GetOperation struct {
	cars domain.CarRepository
}

type GetOperationParams struct {
	ID int64
}

func NewGetOperation(cars domain.CarRepository) *GetOperation {
	return &GetOperation{cars: cars}
}

func (o *GetOperation) Get(ctx context.Context, params *GetOperationParams) (*domain.Car, error) {
	c, err := o.cars.Get(ctx, &domain.GetCarFilter{
		ID: params.ID,
	})
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrCarNotFound
		}

		return nil, err
	}

	return c, nil
}
