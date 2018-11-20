package sets

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager_PutUserSet_GetUserSet(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	m := newManager(t, db)

	tcs := []struct {
		set *v0.UserSet
		err error

		desc string
	}{
		{
			set:  func() *v0.UserSet { set := &v0.UserSet{Set: "testset1"}; return set }(),
			err:  nil,
			desc: "marshallable set should not return an error",
		},
	}

	for _, tc := range tcs {
		tx, err := db.Begin()
		require.NoError(t, err)
		err = m.UserSet.Put(tx, tc.set)
		if tc.err != nil {
			require.EqualError(t, err, tc.err.Error())
			tx.Rollback()
			continue
		}
		require.NoError(t, err)

		set, err := m.UserSet.Get(tx, "testset1")
		require.NoError(t, err, tc.desc)
		assert.EqualValues(t, tc.set.Set, set.Set, tc.desc)
		assert.EqualValues(t, tc.set.Users, set.Users, tc.desc)
	}
}
