package signing

import (
	"io"
)

// SigningService takes a transaction object, reads/unmarshals it's JSON payload,
// signs it as appropriate for the respective digital asset, marshals it back to a []byte,
// and sets the transaction object
type SigningService interface {
	Sign(reader io.Reader) (io.Reader, error)

	// Verify checks if the signature on a transaction is correct.
	Verify(reader io.Reader) (bool, error)
}

// ValidationService takes a signed or unsigned transaction and validates the if the
// transaction is internally logical (no negative amount if the currency does not allow it, etc.)
//
// transaction validity is asset specific.
type ValidationService interface {
	Validate(reader io.Reader) error
}
