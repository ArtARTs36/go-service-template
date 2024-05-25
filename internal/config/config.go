package config

import "time"

type Config struct {
	DB  DBConfig `envPrefix:"DB_"`
	Log Log      `envPrefix:"LOG_"`
}

type DBConfig struct {
	DSN string `env:"DSN,required"`
}

type Log struct {
	Level  string `env:"LEVEL"`
	Sentry struct {
		DSN          string        `env:"DSN"`
		FlushTimeout time.Duration `env:"FLUSH_DURATION" envDefault:"2s"`
	} `envPrefix:"SENTRY_"`
}
