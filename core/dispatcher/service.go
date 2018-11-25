package dispatcher

import (
	"github.com/carapace/core/api/v0/proto"
	"sync"
)

type Provider interface {
	Generate(transaction v0.Transaction, identity v0.Identity) (v0.Transaction, error)
}

type Service struct {
	mu        sync.RWMutex
	providers map[v0.Asset]Provider
}

func Register() {

}
