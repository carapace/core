//go:generate mockgen -destination=mocks/handler_mock.go -package=mock github.com/carapace/core/internal/v0 Handler

package v0_handler

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
)

// Handler defines the common interface for the ownerSet, walletSet and userSet handlers
type Handler interface {
	Handle(ctx context.Context, config v0.Config) (v0.Response, error)
}
