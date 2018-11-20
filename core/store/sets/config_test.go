package sets

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_Add(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	m := newManager(t, db)

	tx, err := db.Begin()
	require.NoError(t, err)

	config := &v0.Config{Header: &v0.Header{Increment: 1}}

	// first time adding config should not error
	err = m.Config.Add(tx, config)
	require.NoError(t, err)

	// second time should, since the increment is the same
	err = m.Config.Add(tx, config)
	require.Error(t, err)

	// increasing the increment should work again
	config.Header.Increment = 2
	err = m.Config.Add(tx, config)
	require.NoError(t, err)
}
