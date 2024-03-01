package domain

import "context"

type CarRepository interface {
	Find(ctx context.Context, id int64) (*Car, error)
}

type Car struct {
	ID int64 `db:"id"`
}
