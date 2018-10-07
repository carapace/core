package asset

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type AssetService interface {
	ValidationService
	Configure(config *Config) error
	Create(transaction v1.Transaction) (*v1.TransactionResponse, error)
}
