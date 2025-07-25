package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseHost string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort uint16 `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"cloud_castle"`
	DatabaseUser string `env:"DATABASE_USER" envDefault:"cloud_castle"`
	DatabasePass string `env:"DATABASE_PASS" envDefault:"topsecret"`
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DatabaseHost, c.DatabasePort, c.DatabaseUser, c.DatabasePass, c.DatabaseName)
}

func (c *Config) BuildDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.DatabaseUser, c.DatabasePass, c.DatabaseHost, c.DatabasePort, c.DatabaseName)
}

func MustReadConfig() Config {
	var config Config
	env.Must(config, env.Parse(&config))
	return config
}
