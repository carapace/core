package ark

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	assert.NotPanics(t, func() { New(Config{}) })
}

func TestService_Sign_and_Validate(t *testing.T) {
	s := newService()

	tx, err := newJSONtx()
	require.NoError(t, err)

	signed, err := s.Sign(bytes.NewReader(tx))

	require.NoError(t, err)
	ok, err := s.Verify(signed)

	require.NoError(t, err)
	assert.True(t, ok)
}
