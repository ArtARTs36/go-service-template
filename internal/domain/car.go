package domain

import "context"

type CarRepository interface {
	Find(ctx context.Context, id int) (*Car, error)
}

type Car struct {
	ID int
}
