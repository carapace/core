package test

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/storage"
	"github.com/ulule/deepcopier"
)

var store *storage.MemStore

func GetStore() storage.SecretsStore {
	s := &storage.MemStore{}
	err := deepcopier.Copy(store).To(s)
	if err != nil {
		panic(err)
	}
	return s
}

func init() {
	store = storage.NewMemstore()

	// preseed the store with keys
	for key := range v1.Asset_name {
		store.PutSecret(GetNamespace(1), v1.Asset(key), GenSecret(v1.Asset(key), GetNamespace(1)))
	}
}
