package userset

import (
	"context"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/suite"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"

	"github.com/carapace/core/api/v0/proto"
)

func TestHandler_ConfigService(t *testing.T) {
	tcs := []struct {
		config  *v0.Config
		userSet *v0.UserSet
		res     *v0.Response

		err  error
		desc string
	}{
		{
			config: &v0.Config{
				Header: &v0.Header{
					ApiVersion: "v0",
					Kind:       "UserSet",
				},
			},
			userSet: &v0.UserSet{
				Set: "robot set",
				Users: []*v0.User{{
					Name:              "roboto",
					Email:             "k.l.kubat@gmail.com",
					PrimaryPublicKey:  []byte("somekeyatrandom"),
					RecoveryPublicKey: []byte("lessrandomseed"),
				},
				},
			},
			err: nil,
			res: &v0.Response{
				Code: v0.Code_OK,
			},

			desc: "simple userset creation should pass",
		},
	}

	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	store, err := core.NewStore(db)
	require.NoError(t, err)
	handler := New(store)

	for _, tc := range tcs {
		any, err := ptypes.MarshalAny(tc.userSet)
		require.NoError(t, err)
		tc.config.Spec = any

		tx, err := handler.store.Begin()
		require.NoError(t, err)

		ctx := core.ContextWithTransaction(context.Background(), tx)

		res, err := handler.ConfigService(ctx, tc.config)
		if tc.err == nil {
			require.NoError(t, err)
		} else {
			assert.EqualError(t, err, tc.err.Error())
		}
		assert.Equal(t, tc.res.Code, res.Code)
	}
}
