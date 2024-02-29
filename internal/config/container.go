package config

import (
	"go-service-template/internal/application/car"
	"log"

	"github.com/jmoiron/sqlx"

	"go-service-template/internal/domain"
	"go-service-template/internal/infrastructure/repository"
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
		Repositories struct {
			CarRepository domain.CarRepository
		}
	}
}

func InitContainer(conf *Config) *Container {
	db, err := sqlx.Connect("postgres", conf.DB.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	cont := &Container{}
	cont.Infrastructure.DB = db

	cont.initRepositories()
	cont.initOperations()

	return cont
}

func (c *Container) initRepositories() {
	c.Infrastructure.Repositories.CarRepository = repository.NewPGCarRepository(c.Infrastructure.DB)
}

func (c *Container) initOperations() {
	c.Application.Operations.Car.Get = car.NewGetOperation(c.Infrastructure.Repositories.CarRepository)
}
