package identity

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

	set, err := h.store.Sets.Identity.All(ctx, tx)
	if err != nil {
		if err != sets.ErrNotExist {
			return nil, err
		}
	}

	res := []*any.Any{}
	for _, msg := range set {
		anyMSG, marshallErr := ptypes.MarshalAny(msg)
		if marshallErr != nil {
			return nil, marshallErr
		}
		res = append(res, anyMSG)
	}

	return &v0.Info{
		ApiVersion:  "v0",
		Mode:        v0.Mode_Debug,
		Semantic:    "v0.0.1",
		Kinds:       []string{IdentitySet},
		Description: "IdentitySets create on-chain entitities (accounts, wallets or smart-contracts",
		Sets:        res,
	}, nil
}
