package ownerset

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core/store/sets"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func (h *Handler) InfoService(ctx context.Context) (*v0.Info, error) {
	tx, err := h.store.Begin(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var setAny *any.Any
	set, err := h.store.Sets.OwnerSet.Get(ctx, tx)
	if err != nil {
		if err == sets.ErrNotExist {
			setAny = nil
		} else {
			return nil, err
		}
	} else {
		setAny, err = ptypes.MarshalAny(set)
		if err != nil {
			return nil, err
		}
	}

	return &v0.Info{
		ApiVersion:  "v0",
		Mode:        v0.Mode_Debug,
		Semantic:    "v0.0.1",
		Kinds:       []string{OwnerSet},
		Description: "ownerSet sets the owners of a carapace node.",
		Sets:        []*any.Any{setAny},
	}, nil
}
