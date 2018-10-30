package router

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"sync"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/internal/core"
	"github.com/pkg/errors"
)

var _ core.Router = &router{} // compile time check to verify we match the core interface

// V0 router implements simple routing, matching apiVersion + kind to handlers using
// a map
type V0 map[string]map[string]interface{ ConfigService(context.Context, *v0.Config) (*v0.Response, error) }

// Route matches a config by Header.ApiVersion and Header.Kind to a CoreServiceServer (a handler)
func (v V0) Route(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	var handler interface{ ConfigService(context.Context, *v0.Config) (*v0.Response, error) }
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
	infohandler interface{ InfoService(context.Context, *empty.Empty) (*v0.Info, error) }
}

// Register allows configuration handlers to self register. It will panic if a handler is registered twice.
func (r *router) Register(apiVersion, kind string, handler interface{ ConfigService(context.Context, *v0.Config) (*v0.Response, error) }) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// initializing the nil map
	if _, ok := r.V0[apiVersion]; !ok {
		r.V0[apiVersion] = make(map[string]interface{ ConfigService(context.Context, *v0.Config) (*v0.Response, error) })
	}

	if _, ok := r.V0[apiVersion][kind]; ok {
		panic("internal/router: attempting to register an already registered apiVersion and kind")
	}
	r.V0[apiVersion][kind] = handler
}

func (r *router) InfoService(ctx context.Context, e *empty.Empty) (info *v0.Info, err error) {
	return r.infohandler.InfoService(ctx, e)
}

func (r *router) RegisterInfoService(service interface{ InfoService(context.Context, *empty.Empty) (*v0.Info, error) }) {
	r.infohandler = service
}


// Router is a global router, which packages may use in their package/install initialization code to self register,
// instead of having to instantiate the router in the main file.
var Router = &router{mu: sync.RWMutex{}, V0: make(map[string]map[string]interface{ ConfigService(context.Context, *v0.Config) (*v0.Response, error) })}
