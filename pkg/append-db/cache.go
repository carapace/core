package append

import (
	"github.com/carapace/core/pkg/t-cache"
)

const (
	chainHash = "chainhash"
	keys      = "keys"
)

// Cache defines an interface around a cache required by append-db
type Cache interface {
	SetChainHash(uint64)
	ChainHash() uint64
	KeyExists(key string) bool
	AddKey(key string)

	Lock() Committer
	Unlock()
}

// Committer defines a return value by cache, to ensure no cache corruption takes place
type Committer interface {
	Commit()
	Rollback()
}

// compile time assertion to ensure the interface matches
var _ Cache = &DefaultCache{}

// DefaultCache fulfilles the Cache interface
//
// It is based on t-cache, just a thin wrapper
type DefaultCache struct {
	*cache.Cache
}

// SetChainHash overrides the previous chainhash value
func (d *DefaultCache) SetChainHash(val uint64) {
	d.Cache.Set(chainHash, val)
}

// ChainHash returns the current chainhash
func (d *DefaultCache) ChainHash() uint64 {
	v := d.Cache.Get(chainHash)
	val, ok := v.(uint64)
	if !ok {
		panic("append-db: cache has been manually edited")
	}
	return val
}

// KeyExists checks if a key is present in the cache
func (d *DefaultCache) KeyExists(key string) bool {
	v := d.Cache.Get(keys)
	ks, ok := v.(map[string]struct{})
	if !ok {
		panic("append-db: cache has been manually edited")
	}

	_, ok = ks[key]
	return ok
}

// AddKey adds a key to the cache
func (d *DefaultCache) AddKey(key string) {
	v := d.Cache.Get(keys)
	ks, ok := v.(map[string]struct{})

	if ks == nil {
		ks = make(map[string]struct{})
	} else {
		if !ok {
			panic("append-db: cache has been manually edited")
		}
	}

	ks[key] = struct{}{}
	d.Cache.Set(keys, ks)
}

// Lock creates a writelock on the cache and returns a Committer
func (d *DefaultCache) Lock() Committer {
	return d.Cache.Lock()
}

// NewDefaultCache is the constructor for cache
func NewDefaultCache() (d *DefaultCache) {
	return &DefaultCache{
		Cache: cache.New(),
	}
}
