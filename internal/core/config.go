package core

import (
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger
	Router Router

	HealthManager

	Health Health
}

type Health struct {
	Port string
	Host string
}

func (c Config) Build() (*Config, error) {

	return &c, nil
}