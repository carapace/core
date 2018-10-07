package storage

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"sync"
)

type MemStore struct {
	mu    *sync.RWMutex
	store map[string]pair
}

type pair struct {
	secret v1.Secret
	usable bool
}

func (m *MemStore) GetSecret(namespace v1.Namespace, asset v1.Asset) (v1.Secret, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	pair := m.store[namespace.String()+asset.String()]
	if !pair.usable {
		return v1.Secret{}, SecretLocked
	}
	return pair.secret, nil
}

func (m *MemStore) PutSecret(namespace v1.Namespace, asset v1.Asset, secret v1.Secret) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[namespace.String()+asset.String()] = pair{secret: secret, usable: true}
	return nil
}

func (m *MemStore) LockSecret(namespace v1.Namespace, asset v1.Asset) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	s := m.store[namespace.String()+asset.String()]
	s.usable = false
	m.store[namespace.String()+asset.String()] = s
	return nil
}

func NewMemstore() *MemStore {
	return &MemStore{
		mu:    &sync.RWMutex{},
		store: make(map[string]pair),
	}
}
