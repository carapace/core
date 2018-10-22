package test

import (
	"github.com/ory/dockertest"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
	"testing"
	"time"
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

// If the machine running the tests is not very beefy, the maximum number of containers deployed can be configured,
// in some cases allowing for faster test execution. By default there is no limit.
type limiter struct {
	mu         *sync.RWMutex
	resourceNr int
}

func init() {
	viper.BindEnv("TEST_RESOURCE_MAX")
	viper.SetDefault("TEST_RESOURCE_MAX", 1000)
}

func (l limiter) allow() (clear func()) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for {
		if l.resourceNr < viper.GetInt("TEST_RESOURCE_MAX") {
			l.resourceNr++
			return l.decrement
		}
		time.Sleep(1 * time.Second)
	}
}

func (l limiter) decrement() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.resourceNr = l.resourceNr - 1
}

var resourceLimit = limiter{mu: &sync.RWMutex{}, resourceNr: 0}
