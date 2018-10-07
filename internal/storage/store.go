package storage

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type SecretsStore interface {
	// GetSecrets returns the secrets specific to a namespace and asset
	GetSecret(namespace v1.Namespace, asset v1.Asset) (v1.Secret, error)

	// Store a secret under a namespace and asset. Order of secrets should remain, so fifo
	PutSecret(namespace v1.Namespace, asset v1.Asset, secret v1.Secret) error
}
