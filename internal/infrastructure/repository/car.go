package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"go-service-template/internal/domain"
)

type PGCarRepository struct {
	db *sqlx.DB
}

func NewPGCarRepository(db *sqlx.DB) domain.CarRepository {
	return &PGCarRepository{
		db: db,
	}
}

func (r *PGCarRepository) Find(ctx context.Context, id int) (*domain.Car, error) {
	var car *domain.Car

	err := r.db.GetContext(ctx, &car, "SELECT * FROM cars WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &domain.Car{
		ID: id,
	}, nil
}
