package config

type Config struct {
	DB  DBConfig  `envPrefix:"DB_"`
	Log LogConfig `envPrefix:"LOG_"`
}

type DBConfig struct {
	DSN string `env:"DSN,required"`
}

type LogConfig struct {
	Level string `env:"LEVEL"`
}
