//go:generate mockgen -source=car_repo.go -package=repository -destination=./../infrastructure/repository/car_mock.go CarRepository
package domain

import (
	"context"
)

type CarRepository interface {
	Get(ctx context.Context, filter *GetCarFilter) (*Car, error)
	List(ctx context.Context, filter *ListCarFilter) ([]*Car, error)
	Create(ctx context.Context, car *Car) (*Car, error)
	Update(ctx context.Context, car *Car) (*Car, error)
	Delete(ctx context.Context, filter *DeleteCarFilter) (count int64, err error)
}

type UserRepository interface {
	Get(ctx context.Context) string
}

type ListCarFilter struct {
	IDs []int64
}

type GetCarFilter struct {
	ID int64
}

type DeleteCarFilter struct {
	IDs []int64
}
