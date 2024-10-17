package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/artarts36/go-service-template/internal/domain"
)

type Group struct {
	CarRepository domain.CarRepository
}

func NewGroup(db *sqlx.DB) *Group {
	return &Group{
		CarRepository: NewPGCarRepository(db),
	}
}
