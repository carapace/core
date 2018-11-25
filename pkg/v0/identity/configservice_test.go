package identity

import (
	"context"
	"github.com/carapace/core/core/store/sets"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"testing"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/core/mocks"
	"github.com/carapace/core/pkg/v0"
	"github.com/golang/mock/gomock"
	"github.com/ory/ladon"
	manager "github.com/ory/ladon/manager/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_ConfigService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store, sCtrl, cleanup := mock.NewStoreMock(t, ctrl)
	defer cleanup()

	handler := New(store, &ladon.Ladon{
		Manager: manager.NewMemoryManager(),
	})

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
						Permissions: []*v0.Permission{
							{
								Conditions: []*v0.Condition{
									{
										Name: v0.ConditionNames_AuthLevelGreater,
										Args: func() *any.Any {
											args := &v0.AuthLevelGreaterArg{Level: 1}
											a, err := ptypes.MarshalAny(args)
											require.NoError(t, err)
											return a
										}(),
									},
								},
							},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			err: errors.New("identityHandler unable to obtain existing identity: oops"),
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(nil, nil),
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
						Permissions: []*v0.Permission{
							{
								Conditions: []*v0.Condition{
									{
										Name: v0.ConditionNames_AuthLevelGreater,
										Args: func() *any.Any {
											args := &v0.AuthLevelGreaterArg{Level: 1}
											a, err := ptypes.MarshalAny(args)
											require.NoError(t, err)
											return a
										}(),
									},
								},
							},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
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
						Permissions: []*v0.Permission{
							{
								Conditions: []*v0.Condition{
									{
										Name: v0.ConditionNames_AuthLevelGreater,
										Args: func() *any.Any {
											args := &v0.AuthLevelGreaterArg{Level: 1}
											a, err := ptypes.MarshalAny(args)
											require.NoError(t, err)
											return a
										}(),
									},
								},
							},
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
					Permissions: []*v0.Permission{},
				}, nil),
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(&v0.User{Name: "Karel", AuthLevel: 9}, nil),
			},
			res: &v0.Response{Code: v0.Code_Forbidden},
		},
		{
			desc: "user having sufficient auth should alter obj.",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Permissions: []*v0.Permission{
							{
								Subjects: []string{"Karel"},
								Actions:  []v0.Action{v0.Action_Alter},
								Effect:   v0.Effect_Allow,
								Conditions: []*v0.Condition{
									{
										Name: v0.ConditionNames_AuthLevelGreater,
										Args: func() *any.Any {
											args := &v0.AuthLevelGreaterArg{Level: 1}
											a, err := ptypes.MarshalAny(args)
											require.NoError(t, err)
											return a
										}(),
									},
								},
							},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(&v0.User{Name: "Karel", AuthLevel: 9}, nil),
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(nil, sets.ErrNotExist),
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
			desc: "alteration causing user lockout should return error",
			config: &v0.Config{
				Header: &v0.Header{Kind: IdentitySet},
				Spec: func() *any.Any {
					id := &v0.Identity{
						Name:  "MySet",
						Asset: v0.Asset_BTC,
						Permissions: []*v0.Permission{
							{
								Subjects: []string{"Karel"},
								Actions:  []v0.Action{v0.Action_Alter},
								Effect:   v0.Effect_Allow,
								Conditions: []*v0.Condition{
									{
										Name: v0.ConditionNames_AuthLevelGreater,
										Args: func() *any.Any {
											args := &v0.AuthLevelGreaterArg{Level: 20}
											a, err := ptypes.MarshalAny(args)
											require.NoError(t, err)
											return a
										}(),
									},
								},
							},
						},
					}
					a, err := ptypes.MarshalAny(id)
					require.NoError(t, err)
					return a
				}(),
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), []byte("key")).Return(&v0.User{Name: "Karel", AuthLevel: 9}, nil),
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), "MySet").Return(nil, sets.ErrNotExist),
			},
			misc: []func(){
				func() {
					sCtrl.DB.ExpectCommit()
				},
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
		assert.Equal(t, tc.res.Code, res.Code, res.Err)
	}
}
