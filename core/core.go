package core

import (
	"sync"
)

// App is the godlevel struct, fulfilling the gRPC interfaces defined by the api.
type App struct {
	*Config
	ready bool
	mu    sync.RWMutex
}

func (a *App) Start() {
	// starts initializations
	a.Config.HealthManager.Start()
	go a.serveHTTP()
	a.ready = true
	// initializations finished, normal operations may commence now
}

type Option func(*App) error

func New(config *Config, options ...Option) (*App, error) {
	app := &App{Config: config}

	for _, opt := range options {
		err := opt(app)
		if err != nil {
			return nil, err
		}
	}

	return app, nil
}
