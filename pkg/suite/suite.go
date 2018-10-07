package test

import (
	"sync"
)

var Suite = New()

type suite struct {
	mu    sync.RWMutex
	ports map[string]struct{}
}

func New() *suite {
	return &suite{
		ports: make(map[string]struct{}),
	}
}
