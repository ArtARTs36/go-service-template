package car

import (
	"context"
	"errors"
	"log/slog"

	"github.com/artarts36/go-service-template/internal/infrastructure/repository"

	"github.com/artarts36/go-service-template/internal/domain"
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
	c, err := o.cars.Find(ctx, params.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrCarNotFound
		}

		slog.
			With(slog.String("err", err.Error())).
			ErrorContext(ctx, "unable to get car")

		return nil, err
	}

	return c, nil
}
