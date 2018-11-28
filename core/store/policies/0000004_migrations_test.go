package policies

import (
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUp0000004(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback()

	err = Up0000004(tx)
	assert.NoError(t, err)
}

func TestDown0000004(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback()

	err = Up0000004(tx)
	require.NoError(t, err)

	err = Down0000004(tx)
	assert.NoError(t, err)
}
