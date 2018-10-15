package core

import (
	"context"

	"github.com/carapace/core/api/v1/proto/generated"
)

// ConfigIngress defines an interface for the controller dealing with new configuration ingress.
// context cancellation is honored up to the point of no return (commit into auditable store),
// cancellation after that point is only fulfilled after the store returns.
// Handlers are activated if the config is committed in store.
type ConfigIngress interface {
	In(context.Context, v1.Config) (result string, err error)
}
