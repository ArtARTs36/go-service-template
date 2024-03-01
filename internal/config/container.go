package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/go-service-template/internal/application/car"
	"github.com/artarts36/go-service-template/internal/domain"
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
		Repositories struct {
			CarRepository domain.CarRepository
		}
	}
}

func InitContainer(conf *Config) *Container {
	cont := &Container{}
	cont.setupLogger(&conf.Log)

	db, err := sqlx.Connect("postgres", conf.DB.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	log.Debug("[container] connected to db")

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
