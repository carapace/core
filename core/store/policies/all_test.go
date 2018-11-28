package policies

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core/store/sets"
	"github.com/carapace/core/pkg/suite"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager_Add_All(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	manager := newManager(t, db)
	tx, err := db.Begin()

	err = sets.New().Identity.Put(context.Background(), tx, &v0.Identity{
		Name: "mywallet",
	})
	require.NoError(t, err)

	err = manager.Add(context.Background(), tx, &v0.Policy{
		Description: "test",
		Actions:     []v0.Action{v0.Action_Alter, v0.Action_Use},
		Effect:      v0.Effect_Allow,
		Conditions:  []*v0.Condition{{Name: v0.ConditionNames_UsersOwns, Args: &any.Any{}}},
		Subjects:    []string{"*"},
	}, "mywallet", "GLOBAL")

	require.NoError(t, err)

	err = manager.Add(context.Background(), tx, &v0.Policy{
		Description: "test2",
		Actions:     []v0.Action{v0.Action_Alter, v0.Action_Use},
		Effect:      v0.Effect_Allow,
		Conditions:  []*v0.Condition{{Name: v0.ConditionNames_UsersOwns, Args: &any.Any{}}},
		Subjects:    []string{"*"},
	}, "mywallet", "GLOBAL")
	require.NoError(t, err)

	policies, err := manager.All(context.Background(), tx)
	require.NoError(t, err)
	assert.Equal(t, 2, len(policies))
}
