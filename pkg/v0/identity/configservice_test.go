package identity

import (
	"context"
	"fmt"
	"testing"

	"github.com/carapace/core/core/store/sets"

	"github.com/carapace/core/pkg/v0"
	"github.com/pkg/errors"

	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_ConfigService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store, sCtrl, cleanup := mock.NewStoreMock(t, ctrl)
	defer cleanup()

	permMock := mock.NewMockPermissionManager(ctrl)
	handler := New(store, permMock)

	tcs := []struct {
		desc     string
		spec     *v0.Identity
		config   *v0.Config
		user     *v0.User
		existing *v0.Identity
		res      *v0.Response
		err      error
		prep     []*gomock.Call
	}{
		{
			desc:   "incorrect header.Kind returns error",
			config: &v0.Config{Header: &v0.Header{Kind: "incorrect"}},
			err:    v0_handler.ErrIncorrectKind,
		},
		{
			desc: "user get error returns error",
			config: &v0.Config{Header: &v0.Header{
				Kind: IdentitySet,
			},
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			spec: &v0.Identity{
				Name:  "mywallet",
				Asset: v0.Asset_BTC,
				Policies: []*v0.Policy{
					{
						ID:          "someID",
						Description: "some desc",
						Conditions: []*v0.Condition{
							{
								Args: mustAnyMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
								Name: v0.ConditionNames_UsersOwns,
							},
						}},
				},
			},
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("oops")),
			},
			err: errors.New("identityHandler unable to obtain signing user: oops"),
		},
		{
			desc: "unknown identity get returns error",
			config: &v0.Config{Header: &v0.Header{
				Kind: IdentitySet,
			},
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			spec: &v0.Identity{
				Name:  "mywallet",
				Asset: v0.Asset_BTC,
				Policies: []*v0.Policy{
					{
						ID:          "someID",
						Description: "some desc",
						Conditions: []*v0.Condition{
							{
								Args: mustAnyMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
								Name: v0.ConditionNames_UsersOwns,
							},
						}},
				},
			},
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("oops")),
			},
			err: errors.New("identityHandler unable to obtain existing identity: oops"),
		},
		{
			desc: "ErrIdentityNotExists creates new identity",
			config: &v0.Config{Header: &v0.Header{
				Kind: IdentitySet,
			},
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			spec: &v0.Identity{
				Name:  "mywallet",
				Asset: v0.Asset_BTC,
				Policies: []*v0.Policy{
					{
						ID:          "someID",
						Description: "some desc",
						Conditions: []*v0.Condition{
							{
								Args: mustAnyMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
								Name: v0.ConditionNames_UsersOwns,
							},
						}},
				},
			},
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(&v0.User{
					Name: "karel",
				}, nil),
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, sets.ErrNotExist),
				permMock.EXPECT().PoliciesAllow(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
				sCtrl.Sets.Identity.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
				permMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
			},
			res: &v0.Response{Code: v0.Code_OK},
		},
		{
			desc: "existing identity will evaluate the existing policies first",
			config: &v0.Config{Header: &v0.Header{
				Kind: IdentitySet,
			},
				Witness: &v0.Witness{Signatures: []*v0.Signature{{Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key")}}}},
			},
			spec: &v0.Identity{
				Name:  "mywallet",
				Asset: v0.Asset_BTC,
				Policies: []*v0.Policy{
					{
						ID:          "someID",
						Description: "some desc",
						Conditions: []*v0.Condition{
							{
								Args: mustAnyMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
								Name: v0.ConditionNames_UsersOwns,
							},
						}},
				},
			},
			prep: []*gomock.Call{
				sCtrl.Users.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(&v0.User{
					Name: "karel",
				}, nil),
				sCtrl.Sets.Identity.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(&v0.Identity{
					Name: "walletset",
				}, nil),
				permMock.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
				permMock.EXPECT().PoliciesAllow(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
				permMock.EXPECT().PoliciesAllow(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
				sCtrl.Sets.Identity.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
				permMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)},
			res: &v0.Response{Code: v0.Code_OK},
		},
	}

	for _, tc := range tcs {
		sCtrl.DB.ExpectBegin()

		if tc.res != nil {
			if tc.res.Code == v0.Code_OK {
				sCtrl.DB.ExpectCommit()
			}
		}

		tx, err := store.Begin(context.Background(), nil)
		require.NoError(t, err)

		if tc.spec != nil {
			tc.config.Spec = mustAnyMarshal(t, tc.spec)
		}

		res, err := handler.ConfigService(core.ContextWithTransaction(context.Background(), tx), tc.config)

		fmt.Println(res)

		if tc.err != nil {
			require.EqualError(t, err, tc.err.Error(), tc.desc)
		} else {
			require.NoError(t, err, tc.desc)
		}

		if tc.res != nil {
			assert.Equal(t, tc.res.Code, res.Code, tc.desc)
		}
	}
}

func mustAnyMarshal(t *testing.T, message proto.Message) *any.Any {
	a, err := ptypes.MarshalAny(message)
	require.NoError(t, err)
	return a
}
