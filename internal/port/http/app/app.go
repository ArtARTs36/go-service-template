package app

import (
	"log/slog"

	"github.com/artarts36/go-service-template/internal/config"
)

type Service struct {
	container *config.Container
}

func New(cfg *Config) (*Service, error) {
	cont, err := config.InitContainer(&cfg.Config)
	if err != nil {
		return nil, err
	}

	return &Service{
		container: cont,
	}, nil
}

func (srv *Service) OnShutdown() {
	if err := srv.container.Infrastructure.DB.Close(); err != nil {
		slog.
			With(slog.String("err", err.Error())).
			Error("[http][service] failed to close db")
	}
}
