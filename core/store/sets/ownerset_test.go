package sets

import (
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager_PutOwnerSet_GetOwnerSet(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	m := newManager(t, db)

	tcs := []struct {
		set *v0.OwnerSet
		err error

		desc string
	}{
		{
			set:  func() *v0.OwnerSet { set := &v0.OwnerSet{}; set.Reset(); return set }(),
			err:  nil,
			desc: "marshallable set should not return an error",
		},
	}

	for _, tc := range tcs {
		tx, err := db.Begin()
		require.NoError(t, err)
		err = m.OwnerSet.Put(tx, tc.set)
		if tc.err != nil {
			fmt.Println(err)
			require.EqualError(t, err, tc.err.Error())
			tx.Rollback()
			continue
		}
		set, err := m.OwnerSet.Get(tx)
		require.NoError(t, err, tc.desc)
		assert.Equal(t, tc.set, set, tc.desc)
	}
}
