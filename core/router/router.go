package router

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"sync"
)

type Router struct {
	mu sync.RWMutex
	V0
}

func New() *Router {
	return &Router{
		mu: sync.RWMutex{},
		V0: make(map[string]map[string]core.APIService),
	}
}

// V0 router implements simple routing, matching apiVersion + kind to handlers using
// a map
type V0 map[string]map[string]core.APIService

// Route matches a config by Header.ApiVersion and Header.Kind to a CoreServiceServer (a handler)
func (v V0) Route(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	var handler interface {
		ConfigService(context.Context, *v0.Config) (*v0.Response, error)
	}
	var ok bool
	if _, ok = v[config.Header.ApiVersion]; !ok {
		return nil, errors.New("unregistered apiVersion")
	}
	if handler, ok = v[config.Header.ApiVersion][config.Header.Kind]; !ok {
		return nil, errors.New("unregistered kind")
	}
	return handler.ConfigService(ctx, config)
}

// Register allows configuration handlers to self register. It will panic if a handler is registered twice.
func (r *Router) Register(handlers ...core.APIService) {
	for _, handler := range handlers {
		r.register(handler)
	}
}

func (r *Router) register(handler core.APIService) {
	r.mu.Lock()
	defer r.mu.Unlock()

	info, err := handler.InfoService()
	if err != nil {
		panic("router.init: " + err.Error())
	}

	// initializing the nil map
	if _, ok := r.V0[info.ApiVersion]; !ok {
		r.V0[info.ApiVersion] = make(map[string]core.APIService)
	}

	for _, kind := range info.Kinds {
		if _, ok := r.V0[info.ApiVersion][kind]; ok {
			panic("router.init: attempting to register an already registered apiVersion and kind")
		}
		r.V0[info.ApiVersion][kind] = handler
	}
}

// InfoService returns each hardcoded info statement on each registered service.
func (r *Router) InfoService(ctx context.Context, e *empty.Empty) (info *v0.RepeatedInfo, err error) {
	info = &v0.RepeatedInfo{}
	info.Reset()
	for _, services := range r.V0 {
		for _, service := range services {
			inf, err := service.InfoService()
			if err != nil {
				return nil, err
			}
			info.Info = append(info.Info, inf)
		}
	}
	return
}
