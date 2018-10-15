package config_store

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core"
	"github.com/carapace/core/internal/scheme"
	"github.com/imdario/mergo"
)

type Store struct {
	Engine
}

func (s *Store) Add(config v1.Config) (core.Response, error) {
	err := s.Put(config)
	if err != nil {
		return nil, err
	}

	all, err := s.GetAll(config.Header.Name)
	if err != nil {
		return nil, err
	}

	merged, err := s.merge(all)
	if err != nil {
		return nil, err
	}
	return s.trigger(*merged)
}

func (s *Store) trigger(conf v1.Config) (core.Response, error) {
	res := make(chan core.Response, 1)
	err := make(chan error, 1)

	go func() {
		handler := scheme.Get(conf.Header.ApiVersion, conf.Header.Kind)
		r, e := handler.Call(conf)
		res <- r
		err <- e
	}()

	return <-res, <-err
}

func (s *Store) merge(confs []*v1.Config) (*v1.Config, error) {
	orignal := confs[0]

	for i := 1; i < len(confs); i++ {
		err := mergo.Merge(&orignal, confs[i], mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	}
	return orignal, nil
}

// Engine defines the storage abstraction, allowing to switch between a relational DB, file or document store if needed.
// The engine uses the required Name field in v1.Config.Header to group configuration items.
type Engine interface {
	Put(v1.Config) error

	// GetAll returns configuration files in ascending order, thus the first item is the first config file
	// of that version-kind-name
	GetAll(name string) ([]*v1.Config, error)
}
