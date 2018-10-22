package cache

import (
	"sync"
)

type Committer struct {
	cache *Cache
	mu    *sync.RWMutex
	store map[string]interface{}
	done  bool
}

func (c *Committer) Commit() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.done {
		panic("already committed the transaction")
	}

	for key, val := range c.store {
		c.cache.cache[key] = val
	}

	c.cache.commiter = nil
	c.cache = nil
	c.done = true
}

func (c *Committer) Rollback() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.done {
		return
	}

	c.cache.commiter = nil
	c.cache = nil
}
