package config

type Config struct {
	DB  DBConfig `envPrefix:"DB_"`
	Log Log      `envPrefix:"LOG_"`
}

type DBConfig struct {
	DSN string `env:"DSN,required"`
}

type Log struct {
	Level string `env:"LEVEL"`
}
