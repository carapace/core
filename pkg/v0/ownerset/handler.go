package ownerset

import (
	"context"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/v0"
	"sync"
)

const (
	// accepted kind header for the ownerSet handler
	OwnerSet = "ownerSet"
)

type Handler struct {
	auth v0_handler.Authenticator

	mu *sync.RWMutex
}

func (h *Handler) ConfigService(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	// Since owners are a unique set in the node, we don't want to have some race condition
	// where two different new ownerSets are concurrently processed. This is mainly a
	// sanity check.
	h.mu.Lock()
	defer h.mu.Unlock()

	if config.Header.Kind != OwnerSet {
		return nil, v0_handler.ErrIncorrectKind
	}

	if !h.auth.HaveOwners() {
		return h.newOwnerSet(ctx, config)
	}

	return h.adjustOwnerSet(ctx, config)
}

func (h *Handler) newOwnerSet(ctx context.Context, config *v0.Config) (*v0.Response, error) {

	set, err := unmarshalAny(config.Spec)
	if err != nil {
		return nil, err
	}

	// first check if all the owners signed the config.
	ok, incorrect, err := h.auth.CheckSignatures(config.Witness)
	if err != nil {
		return nil, err
	}

	if !ok {
		return v0_handler.WriteMSG(v0.Code_BadRequest, fmt.Sprintf("incorrect signature by %s", incorrect)), nil
	}

	// check if all the owners also signed the message
	for _, i := range set.Owners {
		if _, ok := config.Witness.Signatures[i.PrimaryPublicKey]; !ok {
			return v0_handler.WriteMSG(v0.Code_BadRequest, fmt.Sprintf("not all owners signed the set: %s", i.Name)), nil
		}
	}

	err = h.auth.SetOwners(ctx, set)
	return nil, err
}

func (h *Handler) adjustOwnerSet(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	set, err := unmarshalAny(config.Spec)
	if err != nil {
		return nil, err
	}

	root, err := h.auth.GrantRoot(config.Witness)

	if err != nil {

		if err == v0_handler.ErrBackupKeyPresent {
			root, err = h.auth.GrantBackupRoot(config.Witness)
		}

		if err != nil {
			return nil, err
		}
	}

	if !root {
		return v0_handler.WriteMSG(v0.Code_Forbidded, "insufficient quorum for root op"), nil
	}
	err = h.auth.SetOwners(ctx, set)
	return nil, err
}
