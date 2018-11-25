package permissions

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/ory/ladon"
	manager "github.com/ory/ladon/manager/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserAuthGreater(t *testing.T) {
	tcs := []struct {
		desc           string
		user           *v0.User
		conditionLevel int32
		pass           bool
	}{
		{
			desc:           "equal level should pass",
			user:           &v0.User{AuthLevel: 10},
			conditionLevel: 10,
			pass:           true,
		},
		{
			desc:           "greater level should pass",
			user:           &v0.User{AuthLevel: 11},
			conditionLevel: 10,
			pass:           true,
		},
		{
			desc:           "lower level should fail",
			user:           &v0.User{AuthLevel: 9},
			conditionLevel: 10,
			pass:           false,
		},
	}

	condition := AuthLevelGreater{}

	for _, tc := range tcs {
		condition.Level = tc.conditionLevel
		assert.Equal(t, tc.pass, condition.Fulfills(tc.user, nil), "raw call to fulfills")
	}

	warden := &ladon.Ladon{
		Manager: manager.NewMemoryManager(),
	}

	err := warden.Manager.Create(&ladon.DefaultPolicy{
		ID:         "TestUserAuthGreater",
		Resources:  []string{"wallet1"},
		Conditions: ladon.Conditions{"AuthLevelGreater": condition},
		Actions:    []string{"test"},
		Subjects:   []string{"test"},
		Effect:     "allow",
	})
	require.NoError(t, err)

	for _, tc := range tcs {
		condition.Level = tc.conditionLevel
		err = warden.IsAllowed(&ladon.Request{
			Resource: "wallet1",
			Action:   "test",
			Subject:  "test",
			Context: map[string]interface{}{
				"AuthLevelGreater": tc.user,
			},
		})
		if !tc.pass {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)

		}
	}
}
