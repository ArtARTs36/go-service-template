package repository

import (
	"context"

	"go-service-template/internal/domain"

	"github.com/jmoiron/sqlx"
)

type PGCarRepository struct {
	db *sqlx.DB
}

func NewPGCarRepository(db *sqlx.DB) domain.CarRepository {
	return &PGCarRepository{
		db: db,
	}
}

func (r *PGCarRepository) Find(_ context.Context, id int) (*domain.Car, error) {
	var car *domain.Car

	err := r.db.Get(&car, "SELECT * FROM cars WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &domain.Car{
		ID: id,
	}, nil
}
