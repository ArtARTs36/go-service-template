package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/artarts36/go-service-template/internal/domain"
)

const (
	tableCars = "cars"
)

type PGCarRepository struct {
	db *sqlx.DB
}

func NewPGCarRepository(db *sqlx.DB) *PGCarRepository {
	return &PGCarRepository{db: db}
}

func (repo *PGCarRepository) Get(
	ctx context.Context,
	filter *domain.GetCarFilter,
) (*domain.Car, error) {
	var ent domain.Car

	query := goqu.From(tableCars).Select().Limit(1)

	if filter != nil {
		if filter.ID > 0 {
			query = query.Where(goqu.C("id").Eq(filter.ID))
		}
	}

	q, args, err := query.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	err = repo.db.GetContext(ctx, &ent, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return &ent, nil
}

func (repo *PGCarRepository) List(
	ctx context.Context,
	filter *domain.ListCarFilter,
) ([]*domain.Car, error) {
	var ents []*domain.Car

	query := goqu.From(tableCars).Select()

	if filter != nil {
		if len(filter.IDs) > 0 {
			query = query.Where(goqu.C("id").In(filter.IDs))
		}
	}

	q, args, err := query.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	err = repo.db.SelectContext(ctx, &ents, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*domain.Car{}, nil
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return ents, nil
}

func (repo *PGCarRepository) Create(
	ctx context.Context,
	car *domain.Car,
) (*domain.Car, error) {
	query, _, err := goqu.Insert(tableCars).Rows(car).Returning("*").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build insert query: %w", err)
	}

	var created domain.Car

	err = repo.db.GetContext(ctx, &created, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return &created, nil
}

func (repo *PGCarRepository) Update(
	ctx context.Context,
	car *domain.Car,
) (*domain.Car, error) {
	query, _, err := goqu.Update(tableCars).
		Set(car).
		Returning("*").
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	var updated domain.Car

	err = repo.db.GetContext(ctx, &updated, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return &updated, nil
}

func (repo *PGCarRepository) Delete(
	ctx context.Context,
	filter *domain.DeleteCarFilter,
) (count int64, err error) {
	query := goqu.From(tableCars).Delete()

	if filter != nil {
		if len(filter.IDs) > 0 {
			query = query.Where(goqu.C("id").In(filter.IDs))
		}
	}

	q, args, err := query.ToSQL()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	res, err := repo.db.ExecContext(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	count, err = res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get affected rows: %w", err)
	}

	return
}
