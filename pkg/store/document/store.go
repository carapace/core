package document

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core"
	"github.com/imdario/mergo"
)

type Store struct {
	Engine
	encoder Encoder
}

func (s *Store) Add(config v1.Config) (core.Response, error) {
	btes, err := s.encoder.Encode(config)
	if err != nil {
		return nil, err
	}

	err = s.Put(btes)
	if err != nil {
		return nil, err
	}

	all, err := s.GetAll(config.Header.ApiVersion, config.Header.Kind, config.Header.Name)
	if err != nil {
		return nil, err
	}

	res := []*v1.Config{}
	for _, item := range all {
		conf := &v1.Config{}
		err := s.encoder.Decode(conf, item)
		if err != nil {
			return nil, err
		}
		res = append(res, conf)
	}
	_, err = s.merge(res)
	if err != nil {
		return nil, err
	}
	return nil, nil
	// return s.trigger(*merged)
}

func (s *Store) trigger(conf []*v1.Config) (core.Response, error) {
	res := make(chan core.Response, 1)
	err := make(chan error, 1)

	// go func() {
	// 	handler := scheme.Get(conf[0].Header.ApiVersion, conf[0].Header.Kind)
	// 	r, e := handler.Call(context.Background(), conf)
	// 	res <- r
	// 	err <- e
	// }()

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
type Engine interface {
	// Put uses the keys provided, in order, to store the item.
	//
	// empty strings are interpreted as no keys: so the set ["A", "", "C"] is equal to ["A", "C"]
	Put(item []byte, keys ...string) error

	// GetAll returns all items qualified by the key set in ascending order (fifo)
	GetAll(keys ...string) ([][]byte, error)
}
