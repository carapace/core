package mock

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core"
	"github.com/carapace/core/internal/scheme"
)

// Store implements the core.ConfigStore interface. It does not actually store incoming configurations,
// but simply passes the config object to it's handler.
//
// The mock store can be used by using the core.WithMock options.
type Store struct{}

func (s *Store) Add(config v1.Config) (core.Response, error) {
	handler := scheme.Get(config.Header.ApiVersion, config.Header.Kind)
	return handler.Call(config)
}
