package dispatcher

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/ory/ladon"
	"sync"
)

type Provider interface {
	Generate(transaction v0.Transaction, identity v0.Identity) (v0.Transaction, error)
}

type Service struct {
	mu        sync.RWMutex
	providers map[v0.Asset]Provider
	perm      *ladon.Ladon
	store     *core.Store
}
