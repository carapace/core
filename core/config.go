package core

import (
	"github.com/ory/ladon"
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger
	Router Router
	Store  *Store
	Perm   *ladon.Warden

	TXService Dispatcher
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
