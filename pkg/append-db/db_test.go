package append

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("testdb")
	assert.NoError(t, err)
}

func TestDB_Close(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)
	assert.NoError(t, db.Close())
}
