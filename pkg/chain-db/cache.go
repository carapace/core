package append

import (
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrKeyNotExist = errors.New("key not present in cache")
)

type Cache interface {
	GetObjHash(key string) (uint64, error)
	SetObjHash(key string, hash uint64)
	GetChainHash(key string) (uint64, error)
	SetChainHash(key string, hash uint64)
}

type MemCache struct {
	chainmu     *sync.RWMutex
	chainhashes map[string]uint64

	objmu     *sync.RWMutex
	objhashes map[string]uint64
}

func NewMemCace() *MemCache {
	return &MemCache{
		chainmu:     &sync.RWMutex{},
		objmu:       &sync.RWMutex{},
		chainhashes: make(map[string]uint64),
		objhashes:   make(map[string]uint64),
	}
}

func (m *MemCache) GetObjHash(key string) (uint64, error) {
	m.objmu.RLock()
	defer m.objmu.RUnlock()

	if obj, ok := m.objhashes[key]; ok {
		return obj, nil
	}
	return 0, ErrKeyNotExist
}

func (m *MemCache) SetObjHash(key string, hash uint64) {
	m.objmu.Lock()
	defer m.objmu.Unlock()
	m.objhashes[key] = hash
}

func (m *MemCache) GetChainHash(key string) (uint64, error) {
	m.chainmu.RLock()
	defer m.chainmu.RUnlock()

	if obj, ok := m.chainhashes[key]; ok {
		return obj, nil
	}
	return 0, ErrKeyNotExist
}

func (m *MemCache) SetChainHash(key string, hash uint64) {
	m.chainmu.Lock()
	defer m.chainmu.Unlock()
	m.chainhashes[key] = hash
}
