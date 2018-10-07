package asset

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

// ValidationService takes a signed transaction and verifies it has the correct parameters set for
// that specific digital asset
type ValidationService interface {
	Validate(transaction v1.Transaction) error
}
