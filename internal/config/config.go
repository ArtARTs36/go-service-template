package config

type Config struct {
	DB DBConfig `envPrefix:"DB_"`
}

type DBConfig struct {
	DSN string `env:"DSN,required"`
}
