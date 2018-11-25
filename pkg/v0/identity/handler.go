package identity

import (
	"github.com/carapace/core/core"
	"github.com/ory/ladon"
)

const (
	IdentitySet = "IdentitySet"
)

type Handler struct {
	store *core.Store
	perm  *ladon.Ladon
}

func New(store *core.Store, perm *ladon.Ladon) *Handler {
	return &Handler{store: store, perm: perm}
}
