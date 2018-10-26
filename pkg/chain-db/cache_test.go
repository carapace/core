package chaindb

import (
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestMemCache_SetObjHash_GetObjHash(t *testing.T) {
	tcs := []struct {
		key  string
		hash uint64
	}{
		{
			"1",
			1,
		},
		{
			"2",
			100000000,
		},
		{
			"2",
			math.MaxUint64,
		},
	}

	c := NewMemCache()

	for _, tc := range tcs {
		c.SetObjHash(tc.key, tc.hash)
		h, err := c.GetObjHash(tc.key)
		require.NoError(t, err)
		assert.Equal(t, tc.hash, h)
	}
}

func TestMemCache_SetChainHash_GetChainHash(t *testing.T) {
	tcs := []struct {
		key  string
		hash uint64
	}{
		{
			"1",
			1,
		},
		{
			"2",
			100000000,
		},
		{
			"2",
			math.MaxUint64,
		},
	}

	c := NewMemCache()

	for _, tc := range tcs {
		c.SetChainHash(tc.key, tc.hash)
		h, err := c.GetChainHash(tc.key)
		require.NoError(t, err)
		assert.Equal(t, tc.hash, h)
	}
}
