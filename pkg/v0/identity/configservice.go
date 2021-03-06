package identity

import (
	"context"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/core/store/sets"
	"github.com/carapace/core/pkg/responses"
	"github.com/carapace/core/pkg/v0"
	"github.com/golang/protobuf/ptypes"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
)

func (h *Handler) ConfigService(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	if config.Header.Kind != IdentitySet {
		return nil, v0_handler.ErrIncorrectKind
	}

	set := &v0.Identity{}
	err := ptypes.UnmarshalAny(config.Spec, set)
	if err != nil {
		return nil, err
	}

	err = set.Validate()
	if err != nil {
		return response.ValidationErr(err), nil
	}

	tx := core.TXFromContext(ctx)

	signingUser, err := h.store.Users.Get(ctx, tx, config.Witness.Signatures[0].GetPrimaryPublicKey())
	if err != nil {
		return nil, errors.Wrap(err, "identityHandler unable to obtain signing user")
	}

	existing, err := h.store.Sets.Identity.Get(ctx, tx, set.Name)
	if err != nil {
		if err == sets.ErrNotExist {
			return h.newIdentityObj(ctx, set, signingUser, config.Witness)
		}
		return nil, errors.Wrap(err, "identityHandler unable to obtain existing identity")
	}

	// check if we pass existing conditions
	policies, err := h.perm.Get(ctx, tx, existing.Name, "GLOBAL")
	if err != nil {
		return response.Err(err), nil
	}

	err = h.perm.PoliciesAllow(&ladon.Request{
		Subject:  signingUser.Name,
		Resource: set.Name,
		Action:   v0.Action_Alter.String(),
		Context: map[string]interface{}{
			v0.ConditionNames_AuthLevelGTE.String(): signingUser,
			v0.ConditionNames_UsersOwns.String():    signingUser,
			v0.ConditionNames_InSets.String():       signingUser,
			v0.ConditionNames_MultiSig.String():     config.Witness,
		},
	}, policies, "GLOBAL")

	if err != nil {
		return response.PermissionDenied(err.Error()), nil
	}
	return h.newIdentityObj(ctx, set, signingUser, config.Witness)
}

func (h *Handler) newIdentityObj(ctx context.Context, id *v0.Identity, user *v0.User, witness *v0.Witness) (*v0.Response, error) {
	tx := core.TXFromContext(ctx)

	err := h.perm.PoliciesAllow(&ladon.Request{
		Subject:  user.Name,
		Resource: id.Name,
		Action:   v0.Action_Alter.String(),
		Context: map[string]interface{}{
			v0.ConditionNames_AuthLevelGTE.String(): user,
			v0.ConditionNames_UsersOwns.String():    user,
			v0.ConditionNames_InSets.String():       user,
			v0.ConditionNames_MultiSig.String():     witness,
		},
	}, id.Policies, "GLOBAL")

	if err != nil {
		return response.ValidationErr(errors.Wrap(err, "policy would lock out creating user")), nil
	}

	err = h.store.Sets.Identity.Put(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, "identityHandler unable to store Identity")
	}

	err = h.perm.Set(ctx, tx, id.Policies, id.Name, "GLOBAL")
	if err != nil {
		return nil, errors.Wrap(err, "identityHandler unable to store set new policies")
	}
	return response.OK("correctly set Identity"), errors.Wrap(tx.Commit(), "identityHandler commit")
}
