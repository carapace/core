package core

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

// ConfigHandler defines an interface for handlers to handle new configuration states
type ConfigHandler interface {
	Init(Services) error
	Call(v1.Config) (result Response, err error)
}

// Response defines a returned object from a trigger, where all calls block until the trigger is done.
type Response interface {
	Err() error
	MSG() string
}
