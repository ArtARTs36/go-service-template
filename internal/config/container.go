package config

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres

	"github.com/artarts36/go-service-template/internal/application/car"
	"github.com/artarts36/go-service-template/internal/infrastructure/repository"
)

type Container struct {
	Application struct {
		Operations struct {
			Car struct {
				Get *car.GetOperation
			}
		}
	}

	Infrastructure struct {
		DB           *sqlx.DB
		Repositories *repository.Group
	}

	appVersion string
}

func InitContainer(conf *Config, version string) (*Container, error) {
	cont := &Container{
		appVersion: version,
	}
	cont.setupLogger(conf.Log)

	slog.Debug("[container] connecting to db")

	db, err := sqlx.Connect("postgres", conf.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connecting to db: %s", err)
	}

	slog.Debug("[container] connected to db")

	cont.Infrastructure.DB = db

	cont.initRepositories()
	cont.initOperations()

	return cont, nil
}

func (c *Container) initRepositories() {
	c.Infrastructure.Repositories = repository.NewGroup(c.Infrastructure.DB)
}

func (c *Container) initOperations() {
	c.Application.Operations.Car.Get = car.NewGetOperation(c.Infrastructure.Repositories.CarRepository)
}
