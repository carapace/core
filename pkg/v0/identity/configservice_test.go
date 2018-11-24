package identity

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/core/mocks"
	"github.com/carapace/core/core/store/sets"
	"github.com/carapace/core/pkg/v0"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_ConfigService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store, sCtrl, cleanup := mock.NewStoreMock(t, ctrl)
	defer cleanup()

	handler := New(store)

	tcs := []struct {
		desc   string
		err    error
		config *v0.Config
		res    *v0.Response

		prep []*gomock.Call
		misc []func()
	}{
		{
			desc:   "Incorrect kind should return an error",
			err:    v0_handler.ErrIncorrectKind,
			config: &v0.Config{Header: &v0.Header{Kind: "NotIdentitySet"}},
		},
		{
			desc: "Validation error should return Code_BadRequest",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
			},
			res: &v0.Response{Code: v0.Code_BadRequest},
		},
		{
			desc: "Unknown Sets.Identity.Get should return error",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Access: []*v0.AccessProtocol{
							{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
			},
			err: errors.New("identityHandler unable to obtain existing identity: oops"),
			prep: []*gomock.Call{
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(nil, errors.New("oops")),
			},
		},
		{
			desc: "Unknown store.Users.Get error should return error",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Access: []*v0.AccessProtocol{
							{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(nil, nil),
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(nil, errors.New("oops")),
			},
			err: errors.New("identityHandler unable to obtain signing user: oops"),
		},
		{
			desc: "user having insufficient auth should not change the obj.",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Access: []*v0.AccessProtocol{
							{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(&v0.Identity{
					Access: []*v0.AccessProtocol{{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}}},
				}, nil),
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(&v0.User{Name: "Karel", AuthLevel: 9}, nil),
			},
			res: &v0.Response{Code: v0.Code_Forbidden},
		},
		{
			desc: "user having insufficient auth should not change the obj.",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Access: []*v0.AccessProtocol{
							{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(nil, sets.ErrNotExist),
				sCtrl.Sets.Identity.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()),
			},
			misc: []func(){
				func() {
					sCtrl.DB.ExpectCommit()
				},
			},
			res: &v0.Response{Code: v0.Code_OK},
		},
		{
			desc: "user having insufficient auth should not change the obj.",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Access: []*v0.AccessProtocol{
							{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(
					&v0.Identity{
						Asset:  v0.Asset_BTC,
						Access: []*v0.AccessProtocol{{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}}},
					}, nil),
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(&v0.User{Name: "Karel", AuthLevel: 11}, nil),
				sCtrl.Sets.Identity.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
			},
			misc: []func(){
				func() {
					sCtrl.DB.ExpectCommit()
				},
			},
			res: &v0.Response{Code: v0.Code_OK},
		},
		{
			desc: "user having insufficient auth should not change the obj.",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Access: []*v0.AccessProtocol{
							{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 12}},
							{Method: &v0.AccessProtocol_User{User: "Laurens"}},
							{Method: &v0.AccessProtocol_UserSet{UserSet: "Roeiers"}},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(
					&v0.Identity{
						Asset:  v0.Asset_BTC,
						Access: []*v0.AccessProtocol{{Method: &v0.AccessProtocol_AuthLevel{AuthLevel: 10}}},
					}, nil),
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(&v0.User{Name: "Karel", AuthLevel: 11, Set: "klimmers"}, nil),
			},
			res: &v0.Response{Code: v0.Code_BadRequest},
		},
	}

	for _, tc := range tcs {
		sCtrl.DB.ExpectBegin()
		tx, err := store.DB.Begin()
		ctx := core.ContextWithTransaction(context.Background(), tx)
		require.NoError(t, err)

		for _, miscFunc := range tc.misc {
			miscFunc()
		}

		res, err := handler.ConfigService(ctx, tc.config)
		if tc.err != nil {
			assert.EqualError(t, err, tc.err.Error())
			assert.Nil(t, res)
			continue
		}

		require.NoError(t, err)
		assert.Equal(t, tc.res.Code, res.Code)
	}
}
