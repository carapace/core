package userset

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/core/store/sets"
	userstore "github.com/carapace/core/core/store/users"
	"github.com/carapace/core/pkg/responses"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
)

const (
	UserSet = "type.googleapis.com/v0.UserSet"
)

type Handler struct {
	store *core.StoreAPI
}

func New(store *core.StoreAPI) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ConfigService(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	set := v0.UserSet{}
	err := ptypes.UnmarshalAny(config.Spec, &set)
	if err != nil {
		return nil, err
	}

	tx, err := h.store.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = h.store.Sets.UserSet.Put(tx, &set)
	if err != nil {
		return nil, err
	}

	for _, user := range set.Users {
		user.Set = set.Set

		// we don't allow the creation of super users through userSets
		if user.SuperUser {
			return nil, errors.New("creation of superusers through userSets is not allowed")
		}

		// first we check if the user already exists
		usr, err := h.store.Users.Get(tx, user.PrimaryPublicKey)
		if err != nil {
			if err == userstore.ErrUserDoesNotExists {
				err = h.store.Users.Create(tx, *user)
				if err != nil {
					return nil, err
				}
				continue
			}
			return nil, err
		}

		// we don't allow alteration of superusers through userSets.
		if usr.SuperUser {
			return nil, errors.New("creation of superusers through userSets is not allowed")
		}

		err = h.store.Users.Alter(tx, *user)
		if err != nil {
			return nil, err
		}
	}
	return response.OK("correctly created new users"), errors.Wrapf(tx.Commit(), "userSet handler")
}

func (h *Handler) InfoService() (*v0.Info, error) {
	tx, err := h.store.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	st, err := h.store.Sets.UserSet.All(tx)
	if err != nil {
		if err == sets.ErrNotExist {
			st = []*v0.UserSet{}
		}
		return nil, err
	}

	res := []*any.Any{}
	for _, msg := range st {
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
		Kinds:       []string{UserSet},
		Description: "userSets define a group of users",
		Sets:        res,
	}, nil
}
