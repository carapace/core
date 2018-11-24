package sets

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager_PutIdentity_GetIdentity(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	m := newManager(t, db)

	tcs := []struct {
		set *v0.Identity
		err error

		desc string
	}{
		{
			set:  &v0.Identity{Name: "myWallet"},
			err:  nil,
			desc: "marshallable set should not return an error",
		},
	}

	for _, tc := range tcs {
		tx, err := db.Begin()
		require.NoError(t, err)
		err = m.Identity.Put(tx, tc.set)
		if tc.err != nil {
			require.EqualError(t, err, tc.err.Error())
			tx.Rollback()
			continue
		}
		require.NoError(t, err)

		set, err := m.Identity.Get(tx, "myWallet")
		require.NoError(t, err, tc.desc)
		assert.EqualValues(t, tc.set.Name, set.Name, tc.desc)
	}
}
