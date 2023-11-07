package config

import (
	"sync"

	"github.com/ahmadfaizk/goose-app/pkg/env"
)

type (
	Config struct {
		Database DatabaseConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
)

var (
	once sync.Once
	cfg  *Config
)

func Load() *Config {
	once.Do(func() {
		cfg = &Config{
			Database: DatabaseConfig{
				Host:     env.MustGetString("DB_HOST"),
				Port:     env.MustGetInt("DB_PORT"),
				User:     env.MustGetString("DB_USER"),
				Password: env.MustGetString("DB_PASSWORD"),
				Name:     env.MustGetString("DB_NAME"),
			},
		}
	})
	return cfg
}
