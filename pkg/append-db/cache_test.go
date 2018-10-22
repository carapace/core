package append

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultCache_SetChainHash(t *testing.T) {
	c := NewDefaultCache()
	save := c.Lock()

	want := uint64(100000)
	c.SetChainHash(want)
	save.Commit()

	have := c.ChainHash()
	require.Equal(t, want, have)

	save = c.Lock()
	want2 := uint64(1000000)
	c.SetChainHash(want2)
	save.Commit()

	have2 := c.ChainHash()
	require.Equal(t, want2, have2)
}

func TestDefaultCache_AddKey_KeyExists(t *testing.T) {
	c := NewDefaultCache()

	save := c.Lock()
	c.AddKey("mykey")
	save.Commit()

	assert.True(t, c.KeyExists("mykey"))
	assert.False(t, c.KeyExists("non-existant key"))
}
