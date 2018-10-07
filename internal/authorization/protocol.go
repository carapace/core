package authorization

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type Protocol interface {
	// Authorize validates that the accounts together have a high enough authorization to allow the protocol to execute.
	Authorize(protocol v1.Protocol, account ...v1.Account)
}
