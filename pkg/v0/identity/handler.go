package identity

import (
	"github.com/carapace/core/core"
)

const (
	IdentitySet = "IdentitySet"
)

type Handler struct {
	store *core.Store
	perm  core.PermissionManager
}

func New(store *core.Store, perm core.PermissionManager) *Handler {
	return &Handler{store: store, perm: perm}
}
