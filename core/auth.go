//go:generate mockgen -destination=mocks/authenticator_mock.go -package=mock github.com/carapace/core/core Authenticator
//go:generate mockgen -destination=mocks/authorizer_mock.go -package=mock github.com/carapace/core/core Authorizer

package core

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/pkg/errors"
)

var (
	ErrBackupKeyPresent = errors.New("witness contained backup keys")
)

type Authenticator interface {
	// CheckSignatures only checks if the signatures in the witness field match
	// the public keys associated with each signature. It returns an error if it is
	// unable to parse the signatures, the bool indicates if signatures match, the
	// second argument which signature did not match.
	//
	// (the witness field is a map of public keys and signatures)
	// map<string, bytes> Signatures = 1;
	//
	// Check is also responsible for validating the witness field (checking each public key only occurs once, etc.),
	// but shouldn't rely on any backend state.
	Check(message *v0.Config) (bool, string, error)
}

type Authorizer interface {
	// GrantRoot will indicate if root access should be granted. If a backup key
	// is present; ErrBackupKeyPresent is returned
	GrantRoot(ctx context.Context, witness *v0.Witness) (bool, error)

	// GrantBackupRoot will give root access if primary/backup keys are sufficient
	GrantBackupRoot(ctx context.Context, witness *v0.Witness) (bool, error)

	// HaveOwners returns true if the first ownerSet operation has been completed
	//
	// If the node does not have any owners, it should not be able to complete other
	// configuration operations
	HaveOwners(ctx context.Context) (bool, error)

	// GetOwners returns the current owners, or an error if no owners are set or if
	// there is some IO error
	GetOwners(ctx context.Context) (*v0.OwnerSet, error)

	SetOwners(ctx context.Context, set *v0.OwnerSet) error
}
