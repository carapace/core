package auth

import (
	"context"
	"database/sql"
	"github.com/carapace/core/core/store/sets"
	"math"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/pkg/errors"
)

var _ core.Authenticator = &Manager{}
var _ core.Authorizer = &Manager{}

type Manager struct {
	Store *core.Store

	Signer  Signer
	Decoder KeyMarshaller
}

func (m *Manager) GetOwners(ctx context.Context, tx *sql.Tx) (*v0.OwnerSet, error) {
	return m.Store.Sets.OwnerSet.Get(ctx, tx)
}

func (m *Manager) Quorum(ctx context.Context, tx *sql.Tx) (int32, error) {
	set, err := m.GetOwners(ctx, tx)
	if err != nil {
		// just to be sure, let's not return zero, since forgetting
		// to check an error will else always result in root access
		return math.MaxInt32, err
	}
	return set.Quorum, nil
}

func (m *Manager) Check(message *v0.Config) (correct bool, pubkey string, err error) {
	// signatures occur without the witness field set.
	unsigned, err := UnSign(*message)
	if err != nil {
		return false, "", err
	}

	var keyset = make(map[string]struct{})
	for _, signature := range message.Witness.Signatures {

		// Key is a oneof field, we can type switch it, or check it like this
		key := signature.GetPrimaryPublicKey()
		if key == nil {
			key = signature.GetRecoveryPublicKey()
		}

		if key == nil {
			return false, "", errors.New("unset key present in witness")
		}

		if _, ok := keyset[string(key)]; ok {
			return false, "", errors.New("key present multiple times in witness")
		}

		keyset[string(key)] = struct{}{}

		pub, err := m.Decoder.UnmarshalPublic([]byte(key))
		if err != nil {
			return false, "", err
		}

		ok, err := m.Signer.Check(pub, unsigned, signature)
		if err != nil {
			return false, "", err
		}

		if !ok {
			return false, string(key), errors.New("signature does not match")
		}
	}
	return true, "", nil
}

func (m *Manager) HaveOwners(ctx context.Context, tx *sql.Tx) (bool, error) {
	set, err := m.GetOwners(ctx, tx)
	if err != nil {
		if err == sets.ErrNotExist {
			return false, nil
		}
		return false, err
	}
	return set != nil, nil
}
