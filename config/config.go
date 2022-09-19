package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"33306"`
	DBUser     string `env:"DB_USER" envDefault:"sample"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"sample"`
	DBName     string `env:"DB_NAME" envDefault:"sample"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
