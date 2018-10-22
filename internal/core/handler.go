package core

import (
	"context"
	"github.com/carapace/core/api/v1/proto/generated"
)

// ConfigHandler defines an interface for handlers to handle new configuration states
type ConfigHandler interface {
	Init(Services) error

	// call receives all configs with the same apiVersion, kind and name, in ascending order
	Call(context.Context, []*v1.Config, Committer) (result Response, err error)
}

// Response defines a returned object from a trigger, where all calls block until the trigger is done.
type Response interface {
	Err() error
	MSG() string
}

// Committer is used by the handler to relay that the new configuration state has been properly
// handled by calling Commit(). If Rollback() is called, the new configuration state is marked as unhandled
// and placed on a retry queue.
//
// If a handlers context expires, Rollback is called by the initial trigger.
//
// If Commit returns an error, the committer automatically rolls back. Subsequent calls to Rollback or Commit panic.
type Committer interface {
	Commit() error
	Rollback() error
}
