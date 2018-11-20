package core

import (
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger
	Router Router
	Store  *StoreAPI
	HealthManager

	Health Health
}

type Health struct {
	Port string
	Host string
}

func (c Config) Build() (*Config, error) {
	if c.Logger != nil {
		Logger = c.Logger
	}

	return &c, nil
}
