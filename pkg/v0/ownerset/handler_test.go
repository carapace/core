package ownerset

import (
	"context"
	"github.com/carapace/core/pkg/suite"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"testing"

	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/v0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/mock/gomock"

	"github.com/carapace/core/core/mocks"
)

func TestHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store, sCtrl, cleanup := mock.NewStoreMock(t, ctrl)
	defer cleanup()

	authz := mock.NewMockAuthorizer(ctrl)
	handler := New(authz, store)

	tcs := []struct {
		config   *v0.Config
		prep     []*gomock.Call
		err      error
		response *v0.Response
		desc     string
	}{
		{
			prep: []*gomock.Call{
				// no prep, should fail during one of the first few lines
			},
			config: &v0.Config{
				Header: &v0.Header{
					Kind: "incorrectSet"},
				Spec: func() *any.Any {
					set := &v0.OwnerSet{
						Quorum: 1,
						Owners: []*v0.User{{
							Name:              "Karel",
							Email:             "k.l.kubat@gmail.com",
							PrimaryPublicKey:  []byte("123"),
							RecoveryPublicKey: []byte("321")},
						},
					}
					a, err := ptypes.MarshalAny(set)
					require.NoError(t, err)
					return a
				}(),
			},
			err:      v0_handler.ErrIncorrectKind,
			response: nil,
			desc:     "incorrect kind should return v0_handler.ErrIncorrectKind",
		},
		{
			config: &v0.Config{
				Header: &v0.Header{
					Kind: OwnerSet},
				Spec: func() *any.Any {
					set := &v0.OwnerSet{
						Quorum: 1,
						Owners: []*v0.User{{
							Name:              "Karel",
							Email:             "k.l.kubat@gmail.com",
							PrimaryPublicKey:  []byte("123"),
							RecoveryPublicKey: []byte("321")},
						},
					}
					a, err := ptypes.MarshalAny(set)
					require.NoError(t, err)
					return a
				}(),
			},
			err: errors.New("err"),
			prep: []*gomock.Call{
				authz.EXPECT().HaveOwners(gomock.Any(), gomock.Any()).Return(false, errors.New("err")).Times(1),
			},
			response: nil,
			desc:     "error returned by auth.HaveOwners should return the error",
		},
	}

	for _, tc := range tcs {
		sCtrl.DB.ExpectBegin()
		tx, err := store.Begin(context.Background(), nil)
		require.NoError(t, err)
		ctx := core.ContextWithTransaction(context.Background(), tx)

		res, err := handler.ConfigService(ctx, tc.config)
		if err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}
		assert.Equal(t, tc.response, res, tc.desc)
	}
}

func TestHandler_processNewOwners(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store, sCtrl, cleanup := mock.NewStoreMock(t, ctrl)
	defer cleanup()

	handler := Handler{
		store: store,
	}

	tcs := []struct {
		prep   []*gomock.Call
		config *v0.OwnerSet

		err      error
		response *v0.Response
		desc     string
	}{
		{
			prep: []*gomock.Call{
				sCtrl.Sets.OwnerSet.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&v0.OwnerSet{
					Owners: []*v0.User{
						{PrimaryPublicKey: []byte("key")},
					},
				}, nil),
				sCtrl.Users.EXPECT().Delete(gomock.Any(), gomock.Any(), v0.User{
					PrimaryPublicKey: []byte("key"),
				}).Return(nil),
				sCtrl.Sets.OwnerSet.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
				sCtrl.Users.EXPECT().Create(gomock.Any(), gomock.Any(), v0.User{
					PrimaryPublicKey:  []byte("correct key"),
					RecoveryPublicKey: []byte("second correct key"),
					Name:              "Jaap",
					SuperUser:         true}).Return(nil),
			},
			config: &v0.OwnerSet{
				Owners: []*v0.User{
					{
						PrimaryPublicKey:  []byte("correct key"),
						RecoveryPublicKey: []byte("second correct key"),
						Name:              "Jaap"},
				},
			},
			err:      nil,
			response: &v0.Response{Code: v0.Code_OK, MSG: "Correctly altered ownerSet"},
			desc:     "marshalable obj should return no err",
		},
	}

	for _, tc := range tcs {
		sCtrl.DB.ExpectBegin()

		if tc.err == nil {
			sCtrl.DB.ExpectCommit()
		}

		tx, err := handler.store.Begin(context.Background(), nil)
		require.NoError(t, err)
		ctx := core.ContextWithTransaction(context.Background(), tx)

		res, err := handler.alterExisting(ctx, tc.config)
		if tc.err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}
		assert.Equal(t, tc.response.Code, res.Code, tc.desc)
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
	}

	tcs := []struct {
		config *v0.OwnerSet

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config: &v0.OwnerSet{
				Owners: []*v0.User{{
					PrimaryPublicKey:  []byte("correct key"),
					RecoveryPublicKey: []byte("second correct key"),
					Name:              "Jaap",
				},
				},
			},
			err:      nil,
			response: &v0.Response{Code: v0.Code_OK, MSG: "correctly created ownerSet"},
			desc:     "marshalable obj should return no err",
		},
	}

	for _, tc := range tcs {
		tx, err := handler.store.Begin(context.Background(), nil)
		require.NoError(t, err)
		ctx := core.ContextWithTransaction(context.Background(), tx)

		res, err := handler.createNewOwners(ctx, tc.config)
		if tc.err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}
		assert.Equal(t, tc.response.Code, res.Code, tc.desc)
	}
}
