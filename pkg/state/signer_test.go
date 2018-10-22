package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHMACSigner_Sign_Verify(t *testing.T) {
	var (
		key = []byte("mysupersecret")
	)

	tcs := []struct {
		input  interface{}
		output []byte
	}{
		{
			input:  "Laurens' huurcontract",
			output: []byte("5aa9e742c50ed7e5a7732e468207520eb2443c7705ed80947d96e6df7502fad3"),
		},
	}

	h := NewHMAC(key)
	for _, tc := range tcs {
		b, err := h.Sign(tc.input)
		require.NoError(t, err)
		assert.Equal(t, tc.output, b)

		ok, err := h.Verify(tc.input, b)
		require.NoError(t, err)

		assert.True(t, ok)
	}

}
