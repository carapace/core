package store

import (
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	_, err := New(db)
	require.NoError(t, err)
}
