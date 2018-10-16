package test

import (
	"github.com/ory/dockertest"
	"go.uber.org/zap"
)

var pool *dockertest.Pool

func init() {
	Logger.Info("initializing docker pool")

	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		Logger.Panic("Could not connect to docker", zap.Error(err))
	}
	Logger.Info("finished pool initialization")
}
