package ownerset

import (
	"context"
	"github.com/carapace/core/pkg/suite"
	"github.com/pkg/errors"
	"sync"
	"testing"

	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/v0"
	"github.com/golang/protobuf/proto"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/mock/gomock"

	coreMock "github.com/carapace/core/core/mocks"
)

// genAny is used to ignore the error return of marshal.Any when declaring test
// cases
func genAny(t *testing.T, spec proto.Message) *any.Any {
	a, err := ptypes.MarshalAny(spec)
	require.NoError(t, err)
	return a
}

func TestHandler_Handle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	s, err := core.NewStore(db)
	require.NoError(t, err)
	authz := coreMock.NewMockAuthorizer(mockCtrl)
	handler := New(authz, s)

	tcs := []struct {
		config *v0.Config
		prep   []*gomock.Call

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config:   &v0.Config{Header: &v0.Header{Kind: "incorrectSet"}},
			err:      v0_handler.ErrIncorrectKind,
			response: nil,
			desc:     "incorrect kind should return v0_handler.ErrIncorrectKind",
		},
		{
			config: &v0.Config{Header: &v0.Header{Kind: OwnerSet}},
			err:    errors.New("err"),
			prep: []*gomock.Call{
				authz.EXPECT().HaveOwners().Return(false, errors.New("err")).Times(1),
			},
			response: nil,
			desc:     "error returned by auth.HaveOwners should return the error",
		},
	}

	for _, tc := range tcs {
		tx, err := handler.store.Begin()
		require.NoError(t, err)
		res, err := handler.ConfigService(context.Background(), tc.config, tx)
		tx.Rollback()
		if err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}
		assert.Equal(t, tc.response, res, tc.desc)
	}
}

func TestHandler_processNewOwners(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	s, err := core.NewStore(db)
	require.NoError(t, err)

	handler := Handler{
		store: s,
		mu:    &sync.RWMutex{},
	}

	tcs := []struct {
		config *v0.Config

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config: &v0.Config{
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{[]byte("correct key")}}}},
				Spec: genAny(t, &v0.OwnerSet{Owners: []*v0.User{
					{
						PrimaryPublicKey:  []byte("correct key"),
						RecoveryPublicKey: []byte("second correct key"),
						Name:              "Jaap"}}})},
			err:      nil,
			response: &v0.Response{Code: v0.Code_OK, MSG: "Correctly altered ownerSet"},
			desc:     "marshalable obj should return no err",
		},
	}

	for _, tc := range tcs {
		tx, err := handler.store.Begin()
		require.NoError(t, err)
		res, err := handler.processNewOwners(context.Background(), tc.config, tx)
		tx.Rollback()
		if err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}
		assert.Equal(t, tc.response.Code, res.Code, tc.desc)
		break
	}
}

func TestHandler_createNewOwners(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	s, err := core.NewStore(db)
	require.NoError(t, err)

	handler := Handler{
		store: s,
		mu:    &sync.RWMutex{},
	}

	tcs := []struct {
		config *v0.Config

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config: &v0.Config{
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{[]byte("correct key")}}}},
				Spec: genAny(t, &v0.OwnerSet{Owners: []*v0.User{{
					PrimaryPublicKey:  []byte("correct key"),
					RecoveryPublicKey: []byte("second correct key"),
					Name:              "Jaap",
				},
				}})},
			err:      nil,
			response: &v0.Response{Code: v0.Code_OK, MSG: "correctly created ownerSet"},
			desc:     "marshalable obj should return no err",
		},
	}

	for _, tc := range tcs {
		tx, err := handler.store.Begin()
		require.NoError(t, err)
		res, err := handler.createNewOwners(context.Background(), tc.config, tx)
		tx.Rollback()
		if tc.err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}
		assert.Equal(t, tc.response.Code, res.Code, tc.desc)
		break
	}
}
