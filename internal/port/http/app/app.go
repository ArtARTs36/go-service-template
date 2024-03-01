package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/go-service-template/internal/config"
)

type Service struct {
	container *config.Container
}

func New(cfg *Config) *Service {
	cont := config.InitContainer(&cfg.Config)

	return &Service{
		container: cont,
	}
}

func (srv *Service) OnShutdown() {
	if err := srv.container.Infrastructure.DB.Close(); err != nil {
		log.Errorf("[http][service] failed to close db: %s", err.Error())
	}
}
