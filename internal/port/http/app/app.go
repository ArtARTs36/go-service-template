// This file is safe to edit. Once it exists it will not be overwritten

package app

import "github.com/artarts36/go-service-template/internal/config"

type Service struct {
	container *config.Container
}

// New инициализирует сервис
func New(cfg *Config) *Service {
	cont := config.InitContainer(&cfg.Config)

	return &Service{
		container: cont,
	}
}

func (srv *Service) OnShutdown() {
	// do smth on shutdown...
}
