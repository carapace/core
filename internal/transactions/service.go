package transactions

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type AssetService interface {
	v1.TransactionServiceServer
}
