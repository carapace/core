package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotPanics(t, func() {
		New()
	})
}

func TestCacheCommitGet(t *testing.T) {
	c := New()
	save := c.Lock()
	c.Set("key", "value")
	save.Commit()
	c.Unlock()

	val := c.Get("key")
	assert.Equal(t, "value", val)
}

func TestCacheRollbackGet(t *testing.T) {
	c := New()

	save := c.Lock()
	c.Set("key", "value")
	save.Rollback()
	c.Unlock()

	val := c.Get("key")
	assert.Nil(t, val)
}
