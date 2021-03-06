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
			output: []byte{0x26, 0x1b, 0xc7, 0x4d, 0xf5, 0xbc, 0x25, 0x9c, 0x3d, 0x81, 0xbe, 0x90, 0xd9, 0x8a, 0xa3, 0xf7, 0x59, 0xe, 0xf8, 0xc7, 0x8b, 0x1c, 0x31, 0x5d, 0x0, 0x74, 0x39, 0x31, 0xc8, 0xb5, 0xd8, 0xb6},
		},
		{
			input:  "something else",
			output: []byte{0x68, 0x93, 0x3, 0xe6, 0x4c, 0x7f, 0x7f, 0xb3, 0xd6, 0x76, 0x96, 0xb3, 0x50, 0x8d, 0x8c, 0x6e, 0xff, 0xea, 0x20, 0x8f, 0x3e, 0x4f, 0x7, 0xaa, 0x8, 0x37, 0xd0, 0x67, 0x82, 0xd6, 0xf6, 0x8a},
		},
		{
			input:  45678984,
			output: []byte{0xd5, 0x79, 0xb4, 0xac, 0xde, 0x7b, 0x8a, 0x17, 0x82, 0x32, 0xbb, 0xad, 0xa7, 0x7e, 0x3d, 0x53, 0x28, 0xe7, 0xf7, 0x7, 0x61, 0xd3, 0x8b, 0x5f, 0x4a, 0xa2, 0xe5, 0x1f, 0x66, 0x68, 0x62, 0xbc},
		},
		{
			input:  struct{ Key string }{Key: "mykey"},
			output: []byte{0x72, 0x81, 0x49, 0x90, 0xb6, 0xe3, 0xda, 0xea, 0x6, 0x73, 0x1c, 0x4e, 0x49, 0xb1, 0xb1, 0xc7, 0x57, 0x8, 0x1f, 0x3c, 0x26, 0x4c, 0x73, 0xeb, 0x49, 0xa7, 0x47, 0x6c, 0xf5, 0x36, 0x60, 0xeb},
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

func TestHMACSigner_Sign_Verify_Incorrect_Key(t *testing.T) {
	var (
		key  = []byte("mysupersecret")
		key2 = []byte("mysupersecret2")
	)

	tcs := []struct {
		input  interface{}
		output []byte
	}{
		{
			input:  "Laurens' huurcontract",
			output: []byte{0x26, 0x1b, 0xc7, 0x4d, 0xf5, 0xbc, 0x25, 0x9c, 0x3d, 0x81, 0xbe, 0x90, 0xd9, 0x8a, 0xa3, 0xf7, 0x59, 0xe, 0xf8, 0xc7, 0x8b, 0x1c, 0x31, 0x5d, 0x0, 0x74, 0x39, 0x31, 0xc8, 0xb5, 0xd8, 0xb6},
		},
		{
			input:  "something else",
			output: []byte{0x68, 0x93, 0x3, 0xe6, 0x4c, 0x7f, 0x7f, 0xb3, 0xd6, 0x76, 0x96, 0xb3, 0x50, 0x8d, 0x8c, 0x6e, 0xff, 0xea, 0x20, 0x8f, 0x3e, 0x4f, 0x7, 0xaa, 0x8, 0x37, 0xd0, 0x67, 0x82, 0xd6, 0xf6, 0x8a},
		},
		{
			input:  45678984,
			output: []byte{0xd5, 0x79, 0xb4, 0xac, 0xde, 0x7b, 0x8a, 0x17, 0x82, 0x32, 0xbb, 0xad, 0xa7, 0x7e, 0x3d, 0x53, 0x28, 0xe7, 0xf7, 0x7, 0x61, 0xd3, 0x8b, 0x5f, 0x4a, 0xa2, 0xe5, 0x1f, 0x66, 0x68, 0x62, 0xbc},
		},
		{
			input:  struct{ Key string }{Key: "mykey"},
			output: []byte{0x72, 0x81, 0x49, 0x90, 0xb6, 0xe3, 0xda, 0xea, 0x6, 0x73, 0x1c, 0x4e, 0x49, 0xb1, 0xb1, 0xc7, 0x57, 0x8, 0x1f, 0x3c, 0x26, 0x4c, 0x73, 0xeb, 0x49, 0xa7, 0x47, 0x6c, 0xf5, 0x36, 0x60, 0xeb},
		},
	}

	h := NewHMAC(key)
	h2 := NewHMAC(key2)
	for _, tc := range tcs {
		b, err := h.Sign(tc.input)
		require.NoError(t, err)
		assert.Equal(t, tc.output, b)

		ok, err := h2.Verify(tc.input, b)
		require.NoError(t, err)

		assert.False(t, ok)
	}
}
