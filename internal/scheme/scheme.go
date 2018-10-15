package scheme

import (
	"fmt"
	"sync"

	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core"

	"github.com/pkg/errors"
)

var (
	ErrTriggerRegistered = errors.New("scheme: api has already been registered")
)

var s = &Scheme{
	mu:    sync.RWMutex{},
	types: make(map[string]map[string]*registration),
}

type Scheme struct {
	mu sync.RWMutex

	// types is a map of versions, kinds, and trigger hooks
	types map[string]map[string]*registration
}

func Register(version, kind string, hand core.ConfigHandler, validator Validator) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.types[version]; !ok {
		s.types[version] = make(map[string]*registration)
	}
	if _, ok := s.types[version][kind]; ok {
		panic(errors.Wrap(ErrTriggerRegistered, fmt.Sprintf("version: %s, kind: %s", version, kind)))
	}
	s.types[version][kind] = &registration{Handler: hand, Validator: validator}
}

func IsRegistered(version, kind string) bool {
	_, ok := s.Types()[version][kind]
	return ok
}

func Validate(conf v1.Config) error {
	if !IsRegistered(conf.Header.ApiVersion, conf.Header.Kind) {
		return errors.New(fmt.Sprintf("scheme: unregistered apiVersion + kind: %s - %s", conf.Header.ApiVersion, conf.Header.Kind))
	}

	return s.types[conf.Header.ApiVersion][conf.Header.Kind].Validator(conf)
}

func Get(version, kind string) core.ConfigHandler {
	return s.types[version][kind].Handler
}

func (s *Scheme) Types() map[string]map[string]*registration {
	return s.types
}

type registration struct {
	Validator Validator
	Handler   core.ConfigHandler
}
