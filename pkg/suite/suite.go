package test

import (
	"github.com/ory/dockertest"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"testing"
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

const (

	// A unit test requires no outside services
	Unit = 0

	// LocalIntegration tests require DB's and other services, which can be deployed locally
	LocalIntegration = 1

	// Integration tests require services outside of our control
	Integration = 2

	// E2E tests are long, expensive end to end tests
	E2E = 3
)

// Setting up viper to parse flags and environment variables
func init() {
	viper.BindEnv("TEST_LEVEL")
	viper.SetDefault("TEST_LEVEL", 1)
}

// Strategy parses environment variables and flags, and depending on global presets,
// skips tests.
func Strategy(t *testing.T, level int) {
	if viper.GetInt("TEST_LEVEL") < level {
		t.Skip("skipping test, level too high")
	}
}
