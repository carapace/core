//go:generate mockgen -destination=mocks/authentication_mock.go -package=mock github.com/carapace/core/internal/v0 Authenticator

package v0_handler

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
	CheckSignatures(witness *v0.Witness) (bool, string, error)

	// Set new owners for the node; this is a stateful operation and should store
	// the object in some permanent storage
	SetOwners(context.Context, *v0.OwnerSet) error

	// HaveOwners returns if the first ownerSet operation has been completed
	//
	// If the node does not have any owners, it should not be able to complete other
	// configuration operations
	HaveOwners() bool

	// GetOwners returns the current owners, or an error if no owners are set or if
	// there is some IO error
	GetOwners() (*v0.OwnerSet, error)

	// Quorum returns the total owner weight needed to make a rootlevel operation, or
	// an error due to IO or if the owners are not set.
	Quorum() (int32, error)

	// GrantRoot will give root access if primary keys are sufficient. If a backup key
	// is used; ErrBackupKeyPresent is returned
	GrantRoot(witness *v0.Witness) (bool, error)

	// GrantBackupRoot will give root access if primary/backup keys are sufficient
	GrantBackupRoot(witness *v0.Witness) (bool, error)

	// weight returnes the total autorization weight of the witness field. It errors if
	// an owner is present in the field or an unknown public key.
	Weight(witness *v0.Witness) (int32, error)
}
