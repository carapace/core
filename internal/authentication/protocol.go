package authentication

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type Protocol interface {
	// Authenticate validates that the payload was signed by all signatorees
	Authenticate(payload []byte, signatures v1.Signatures) (bool, error)
}
