package chaindb

import (
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrKeyNotExist = errors.New("key not present in cache")
)

// Cache defines the interface for all cacheable operations by chain-db. The cache implementation should ensure
// the values cannot be altered by a different caller, as this would lead to db corruption.
type Cache interface {
	GetObjHash(key string) (uint64, error)
	SetObjHash(key string, hash uint64)
	GetChainHash(key string) (uint64, error)
	SetChainHash(key string, hash uint64)
}

// MemCache implements interface Cache using go maps.
type MemCache struct {
	chainmu     *sync.RWMutex
	chainhashes map[string]uint64

	objmu     *sync.RWMutex
	objhashes map[string]uint64
}

// NewMemCache is the constructor for MemCache
func NewMemCache() *MemCache {
	return &MemCache{
		chainmu:     &sync.RWMutex{},
		objmu:       &sync.RWMutex{},
		chainhashes: make(map[string]uint64),
		objhashes:   make(map[string]uint64),
	}
}

// GetObjHash returns the object hash from the cache, or ErrKeyNotExist
func (m *MemCache) GetObjHash(key string) (uint64, error) {
	m.objmu.RLock()
	defer m.objmu.RUnlock()

	if obj, ok := m.objhashes[key]; ok {
		return obj, nil
	}
	return 0, ErrKeyNotExist
}

// SetObjHash sets/overrides the object hash for a given key
func (m *MemCache) SetObjHash(key string, hash uint64) {
	m.objmu.Lock()
	defer m.objmu.Unlock()
	m.objhashes[key] = hash
}

// GetChainHash returns the chain hash from the cache, or ErrKeyNotExist
func (m *MemCache) GetChainHash(key string) (uint64, error) {
	m.chainmu.RLock()
	defer m.chainmu.RUnlock()

	if obj, ok := m.chainhashes[key]; ok {
		return obj, nil
	}
	return 0, ErrKeyNotExist
}

// SetChainHash sets/overrides the chain hash for a given key
func (m *MemCache) SetChainHash(key string, hash uint64) {
	m.chainmu.Lock()
	defer m.chainmu.Unlock()
	m.chainhashes[key] = hash
}
