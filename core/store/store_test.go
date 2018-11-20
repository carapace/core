package store

import (
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew_AutoMigrate(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	store := New(db)

	require.NoError(t, store.AutoMigrate())
}
