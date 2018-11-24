package identity

import (
	"context"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/core/store/sets"
	"github.com/carapace/core/pkg/responses"
	"github.com/carapace/core/pkg/v0"
	"github.com/golang/protobuf/ptypes"
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

	existing, err := h.store.Sets.Identity.Get(tx, set.Name)
	if err != nil {
		if err == sets.ErrNotExist {
			return h.newIdentityObj(ctx, set)
		}
		return nil, errors.Wrap(err, "identityHandler unable to obtain existing identity")
	}

	signingUser, err := h.store.Users.Get(tx, config.Witness.Signatures[0].GetPrimaryPublicKey())
	if err != nil {
		return nil, errors.Wrap(err, "identityHandler unable to obtain signing user")
	}

	for _, accessAuth := range existing.Access {
		switch accessAuth.Method.(type) {

		case *v0.AccessProtocol_AuthLevel:
			// check if the user is authorized to alter the Identity
			if signingUser.AuthLevel > accessAuth.GetAuthLevel() {
				return h.alterIdentityObj(ctx, set, signingUser, existing)
			}

		case *v0.AccessProtocol_User:
			if signingUser.Name == accessAuth.GetUser() {
				return h.alterIdentityObj(ctx, set, signingUser, existing)
			}

		case *v0.AccessProtocol_UserSet:
			if signingUser.Set == accessAuth.GetUserSet() {
				return h.alterIdentityObj(ctx, set, signingUser, existing)
			}
		}
	}
	return response.MSG(v0.Code_Forbidden, fmt.Sprintf("%s not allowed to alter identity: %s", signingUser.Name, set.Name)), nil
}

func (h *Handler) newIdentityObj(ctx context.Context, id *v0.Identity) (*v0.Response, error) {
	tx := core.TXFromContext(ctx)

	err := h.store.Sets.Identity.Put(tx, id)
	if err != nil {
		return nil, errors.Wrap(err, "identityHandler unable to store Identity")
	}
	return response.OK("correctly set Identity"), errors.Wrap(tx.Commit(), "identityHandler commit")
}

func (h *Handler) alterIdentityObj(ctx context.Context, id *v0.Identity, signingUser *v0.User, existingID *v0.Identity) (*v0.Response, error) {

	if existingID.Asset != id.Asset {
		return response.MSG(v0.Code_BadRequest, "not possible to alter an identities Asset"), nil
	}

	// check if the user is not locking themselves out if the Identity in some form
	if signingUser.AuthLevel > GetMaxAuthFromAccess(id.Access) {
		return h.newIdentityObj(ctx, id)
	}

	if _, ok := GetUserAccessSet(id.Access)[signingUser.Name]; ok {
		return h.newIdentityObj(ctx, id)
	}

	if _, ok := GetUserSetAccessSet(id.Access)[signingUser.Set]; ok {
		return h.newIdentityObj(ctx, id)
	}

	return response.MSG(v0.Code_BadRequest, "not possible to alter an identity; would lead to lockout"), nil
}
