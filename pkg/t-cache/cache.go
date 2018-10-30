package cache

import (
	"sync"
)

// Cache implements a transactional Cache, where the transaction is obtained by calling Lock()
// and released through the returned committer obj.
type Cache struct {
	mu       *sync.RWMutex
	cache    map[string]interface{}
	commiter *Committer
}

// Set creates a new cache item, but does not commit it. Ensure a lock is required before setting values,
// then use the committer to permanently store values
func (c *Cache) Set(key string, val interface{}) {
	c.commiter.mu.Lock()
	defer c.commiter.mu.Unlock()

	c.commiter.store[key] = val
}

// Get returns a cache item, or nil if the key is not present. Get is thread safe.
func (c *Cache) Get(key string) (val interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if val, ok := c.cache[key]; ok {
		return val
	}
	return nil
}

// Lock creates a read lock on the cache, and returns a committer which allows for rolling back the transaction, or
// committing it. Unlock still needs to be called after committing the transaction.
func (c *Cache) Lock() *Committer {
	c.mu.Lock()
	c.commiter = &Committer{cache: c, mu: &sync.RWMutex{}, store: make(map[string]interface{}), done: false}
	return c.commiter
}

// Unlock unlocks the global cache lock. It should always be deferred
func (c *Cache) Unlock() {
	defer c.mu.Unlock()

	if c.commiter != nil {
		panic("t-cache: transaction has not been committed or rolled back")
	}
}

// New is the constructor for the cache
func New() *Cache {
	return &Cache{
		mu:       &sync.RWMutex{},
		cache:    make(map[string]interface{}),
		commiter: nil,
	}
}
