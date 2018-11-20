package sets

import (
	"database/sql"
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// NewManager returns a manager with a migrated DB
func newManager(t *testing.T, db *sql.DB) *Manager {
	m := &Manager{}
	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Commit()

	err = m.AutoMigrate(tx)
	require.NoError(t, err)

	return m
}

func TestManager_AutoMigrate(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	tx, err := db.Begin()
	defer tx.Rollback()
	require.NoError(t, err)

	m := &Manager{}
	err = m.AutoMigrate(tx)
	assert.NoError(t, err)
}
