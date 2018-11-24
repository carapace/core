package identity

import (
	"github.com/carapace/core/core"
)

const (
	IdentitySet = "IdentitySet"
)

type Handler struct {
	store *core.Store
}

func New(store *core.Store) *Handler {
	return &Handler{store: store}
}
