package sets

import (
	"database/sql"
	"github.com/carapace/core/pkg/suite"
	"github.com/pressly/goose"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// NewManager returns a manager with a migrated DB
func newManager(t *testing.T, db *sql.DB) *Manager {
	m := &Manager{
		Config:   &Config{},
		UserSet:  &UserSet{},
		OwnerSet: &OwnerSet{},
	}
	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Commit()

	require.NoError(t, goose.SetDialect("sqlite3"))
	require.NoError(t, goose.Up(db, "."))

	return m
}

func TestInitialUp(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	tx, err := db.Begin()
	defer tx.Rollback()
	require.NoError(t, err)

	err = Up0000001(tx)
	assert.NoError(t, err)
}
