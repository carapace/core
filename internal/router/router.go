package router

import (
	"context"
	"sync"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/internal/core"
	"github.com/pkg/errors"
)

var _ core.Router = V0{} // compile time check to verify we match the core interface

// V0 router implements simple routing, matching apiVersion + kind to handlers using
// a map
type V0 map[string]map[string]v0.CoreServiceServer

// Route matches a config by Header.ApiVersion and Header.Kind to a CoreServiceServer (a handler)
func (v V0) Route(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	var handler v0.CoreServiceServer
	var ok bool
	if _, ok = v[config.Header.ApiVersion]; !ok {
		return nil, errors.New("unregistered apiVersion")
	}
	if handler, ok = v[config.Header.ApiVersion][config.Header.Kind]; !ok {
		return nil, errors.New("unregistered kind")
	}
	return handler.ConfigService(ctx, config)
}

type router struct {
	mu sync.RWMutex
	V0
}

// Register allows configuration handlers to self register. It will panic if a handler is registered twice.
func (r router) Register(apiVersion, kind string, handler v0.CoreServiceServer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// initializing the nil map
	if _, ok := r.V0[apiVersion]; !ok {
		r.V0[apiVersion] = make(map[string]v0.CoreServiceServer)
	}

	if _, ok := r.V0[apiVersion][kind]; ok {
		panic("internal/router: attempting to register an already registered apiVersion and kind")
	}
	r.V0[apiVersion][kind] = handler
}

// Router is a global router, which packages may use in their package/install initialization code to self register,
// instead of having to instantiate the router in the main file.
var Router = &router{mu: sync.RWMutex{}, V0: make(map[string]map[string]v0.CoreServiceServer)}
