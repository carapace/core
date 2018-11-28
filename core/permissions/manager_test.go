package permissions

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/suite"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func mustMarshal(t *testing.T, message proto.Message) *any.Any {
	serialized, err := ptypes.MarshalAny(message)
	require.NoError(t, err)
	return serialized
}

// TestManager_IsAllowed technically tests many functionalites; storing sets and policies, getting policies and deriving
// conditions, and validating requests. Never a bad idea to test too much though.
func TestManager_IsAllowed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	store, err := core.NewStore(db)
	require.NoError(t, err)
	manager := New(store)

	// setup requirements
	wallet := &v0.Identity{
		Name:  "holidayWallet",
		Asset: v0.Asset_BTC,
	}

	tcs := []struct {
		desc     string
		err      error
		policies []*v0.Policy
		request  *ladon.Request
	}{
		{
			desc: "simple permission based on username should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Karel",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_UsersOwns.String(): &v0.User{Name: "Karel"},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet allow",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Karel"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission based on username should fail",
			err:  errors.New("reques"),
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Karel",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_UsersOwns.String(): &v0.User{Name: "Karel"},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet allow",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Karel"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Raka"}}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission based on authlevel should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Karel",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet allow",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Karel"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 9}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission based on authlevel should fail",
			err:  errors.New(""),
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Karel",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet allow",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Karel"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 11}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission based on userset should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Karel",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String(): &v0.User{Set: "boss"},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet allow",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Karel"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_InSets,
							Args: mustMarshal(t, &v0.InSetsArg{Sets: []string{"boss"}}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission based on two multisig should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Karel",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_MultiSig.String(): &v0.Witness{
						Signatures: []*v0.Signature{
							{
								Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key 1")},
							},
							{
								Key: &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: []byte("key 2")},
							},
						},
					},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet allow",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Karel"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_MultiSig,
							Args: mustMarshal(t, &v0.MultiSigArg{PrimaryPublicKeys: [][]byte{
								[]byte("key 1"),
								[]byte("key 2"),
							},
							}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission failing username and userset but passing authlevel should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Raka",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String():       &v0.User{Set: "employee"},
					v0.ConditionNames_UsersOwns.String():    &v0.User{Name: "Raka"},
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet authlevel specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 9}),
						},
					},
				},
				{
					ID:          "holidayWallet ownership specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
						},
					},
				},
				{
					ID:          "holidayWallet ownerset specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_InSets,
							Args: mustMarshal(t, &v0.InSetsArg{Sets: []string{"boss"}}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission failing authlevel and userset but passing username should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Raka",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String():       &v0.User{Set: "employee"},
					v0.ConditionNames_UsersOwns.String():    &v0.User{Name: "Raka"},
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet authlevel specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 11}),
						},
					},
				},
				{
					ID:          "holidayWallet ownership specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Raka"}}),
						},
					},
				},
				{
					ID:          "holidayWallet ownerset specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_InSets,
							Args: mustMarshal(t, &v0.InSetsArg{Sets: []string{"boss"}}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission failing authlevel and username but passing userset should pass",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Raka",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String():       &v0.User{Set: "boss"},
					v0.ConditionNames_UsersOwns.String():    &v0.User{Name: "Karel"},
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet authlevl specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 11}),
						},
					},
				},
				{
					ID:          "holidayWallet ownership specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Raka"}}),
						},
					},
				},
				{
					ID:          "holidayWallet ownerset specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_InSets,
							Args: mustMarshal(t, &v0.InSetsArg{Sets: []string{"boss"}}),
						},
					},
				},
			},
		},
		{
			desc: "simple permission failing authlevel and userset and username should fail",
			err:  errors.New(""),
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Raka",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String():       &v0.User{Set: "employee"},
					v0.ConditionNames_UsersOwns.String():    &v0.User{Name: "Karel"},
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet authlevl specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 11}),
						},
					},
				},
				{
					ID:          "holidayWallet ownership specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Raka"}}),
						},
					},
				},
				{
					ID:          "holidayWallet ownerset specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_InSets,
							Args: mustMarshal(t, &v0.InSetsArg{Sets: []string{"boss"}}),
						},
					},
				},
			},
		},
		{
			desc: "multiconditional policy based permissions should pass if one of the policies passes",
			err:  nil,
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Raka",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String():       &v0.User{Set: "employee"},
					v0.ConditionNames_UsersOwns.String():    &v0.User{Name: "Raka"},
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet authlevel + username specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 11}),
						},
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
						},
					},
				},
				{
					ID:          "holidayWallet ownership specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Raka"}}),
						},
					},
				},
			},
		},
		{
			desc: "multiconditional policy based permissions should fail if all policies fail",
			err:  errors.New(""),
			request: &ladon.Request{
				Resource: "holidayWallet",
				Subject:  "Raka",
				Action:   v0.Action_Alter.String(),
				Context: ladon.Context{
					v0.ConditionNames_InSets.String():       &v0.User{Set: "employee"},
					v0.ConditionNames_UsersOwns.String():    &v0.User{Name: "Raka"},
					v0.ConditionNames_AuthLevelGTE.String(): &v0.User{AuthLevel: 10},
				},
			},

			policies: []*v0.Policy{
				{
					ID:          "holidayWallet authlevel + username specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_AuthLevelGTE,
							Args: mustMarshal(t, &v0.AuthLevelGreaterArg{Level: 11}),
						},
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Karel"}}),
						},
					},
				},
				{
					ID:          "holidayWallet ownership specification",
					Description: "description",
					Actions:     []v0.Action{v0.Action_Alter},
					Effect:      v0.Effect_Allow,
					Subjects:    []string{"Raka"},
					Conditions: []*v0.Condition{
						{
							Name: v0.ConditionNames_UsersOwns,
							Args: mustMarshal(t, &v0.UserOwnsArg{Names: []string{"Laurens"}}),
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		tx, err := store.Begin(context.Background(), nil)
		require.NoError(t, err)

		wallet.Policies = tc.policies

		err = store.Sets.Identity.Put(context.Background(), tx, wallet)
		require.NoError(t, err)

		err = manager.Set(context.Background(), tx, wallet.Policies, wallet.Name, "GLOBAL")
		require.NoError(t, err)

		err = manager.IsAllowed(context.Background(), tx, "GLOBAL", wallet.Name, tc.request)
		if tc.err == nil {
			require.NoError(t, err, tc.desc)
		} else {
			assert.Error(t, err, tc.desc)

		}
		require.NoError(t, tx.Rollback())
	}
}
