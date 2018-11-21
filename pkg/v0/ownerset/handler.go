package ownerset

import (
	"context"
	"database/sql"
	"github.com/carapace/core/pkg/responses"
	"sync"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/v0"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

const (
	// accepted kind header for the ownerSet handler
	OwnerSet = "type.googleapis.com/v0.OwnerSet"
)

var (
	ErrMissingPublicKey = errors.New("missing public key")
)

var _ core.APIService = &Handler{}

type Handler struct {
	authz core.Authorizer
	store *core.Store

	mu *sync.RWMutex
}

func New(authz core.Authorizer, store *core.Store) *Handler {
	return &Handler{
		mu:    &sync.RWMutex{},
		authz: authz,
		store: store,
	}
}

func (h *Handler) ConfigService(ctx context.Context, config *v0.Config, tx *sql.Tx) (*v0.Response, error) {
	// Since owners are a unique set in the node, we don't want to have some race condition
	// where two different new ownerSets are concurrently processed. This is mainly a
	// sanity check.
	h.mu.Lock()
	defer h.mu.Unlock()

	// handling functions should commit the TX; so the error of this defer is ignored.
	defer tx.Rollback()

	if config.Header.Kind != OwnerSet {
		return nil, v0_handler.ErrIncorrectKind
	}

	set := &v0.OwnerSet{}
	err := ptypes.UnmarshalAny(config.Spec, set)
	if err != nil {
		return nil, err
	}

	err = set.Validate()
	if err != nil {
		return response.ValidationErr(err), nil
	}

	have, err := h.authz.HaveOwners()
	if err != nil {
		return nil, err
	}

	if have {
		return h.processNewOwners(ctx, set, tx)
	}
	return h.createNewOwners(ctx, set, tx)
}

func (h *Handler) processNewOwners(ctx context.Context, set *v0.OwnerSet, tx *sql.Tx) (*v0.Response, error) {
	currentSet, err := h.store.Sets.OwnerSet.Get(tx)
	if err != nil {
		return nil, err
	}

	for _, owner := range currentSet.Owners {
		err = h.store.Users.Delete(tx, *owner)
		if err != nil {
			return nil, err
		}
	}
	return h.createNewOwners(ctx, set, tx)
}

func (h *Handler) createNewOwners(ctx context.Context, set *v0.OwnerSet, tx *sql.Tx) (*v0.Response, error) {
	err := h.store.Sets.OwnerSet.Put(tx, set)
	if err != nil {
		return nil, err
	}

	for _, user := range set.Owners {
		err = h.store.Users.Create(tx, *user)
		if err != nil {
			return nil, err
		}
	}
	return v0_handler.WriteSuccess("correctly created ownerSet"), errors.Wrapf(tx.Commit(), "OwnerSet handler.createNewOwners")
}
