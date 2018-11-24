package ownerset

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/responses"
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
}

func New(authz core.Authorizer, store *core.Store) *Handler {
	return &Handler{
		authz: authz,
		store: store,
	}
}

func (h *Handler) ConfigService(ctx context.Context, config *v0.Config) (*v0.Response, error) {
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

	have, err := h.authz.HaveOwners(ctx)
	if err != nil {
		return nil, err
	}

	if have {
		return h.alterExisting(ctx, set)
	}
	return h.createNewOwners(ctx, set)
}

func (h *Handler) alterExisting(ctx context.Context, set *v0.OwnerSet) (*v0.Response, error) {
	tx := core.TXFromContext(ctx)
	currentSet, err := h.store.Sets.OwnerSet.Get(ctx, tx)
	if err != nil {
		return nil, err
	}

	for _, owner := range currentSet.Owners {
		err = h.store.Users.Delete(ctx, tx, *owner)
		if err != nil {
			return nil, err
		}
	}
	return h.createNewOwners(ctx, set)
}

func (h *Handler) createNewOwners(ctx context.Context, set *v0.OwnerSet) (*v0.Response, error) {
	tx := core.TXFromContext(ctx)

	err := h.store.Sets.OwnerSet.Put(ctx, tx, set)
	if err != nil {
		return nil, err
	}

	for _, user := range set.Owners {
		user.SuperUser = true
		err = h.store.Users.Create(ctx, tx, *user)
		if err != nil {
			return nil, err
		}
	}
	return response.OK("correctly created ownerSet"), errors.Wrapf(tx.Commit(), "OwnerSet handler.createNewOwners")
}
