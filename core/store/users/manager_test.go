package user

import (
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/suite"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// NewManager returns a manager with a migrated DB
func newManager(t *testing.T, db *sql.DB) *Manager {
	m := &Manager{}
	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Commit()

	err = m.AutoMigrate(tx)
	require.NoError(t, err)

	return m
}

func TestManager_AutoMigrate(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	manager := &Manager{}
	tx, err := db.Begin()
	require.NoError(t, err)

	err = manager.AutoMigrate(tx)
	assert.NoError(t, err)
}

func TestManager_Create(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	manager := newManager(t, db)

	tx, err := db.Begin()
	require.NoError(t, err)
	err = manager.Create(tx, v0.User{
		PrimaryPublicKey:  []byte("dfghjkjhgrtyuiuytrfvbhy"),
		RecoveryPublicKey: []byte("fghjkjhgfdfghjhgffghjkjhg"),
	})
	assert.NoError(t, err)
}

func TestManager_Create_twice_fails(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	manager := newManager(t, db)

	tx, err := db.Begin()
	require.NoError(t, err)
	err = manager.Create(tx, v0.User{
		PrimaryPublicKey:  []byte("dfghjkjhgrtyuiuytrfvbhy"),
		RecoveryPublicKey: []byte("fghjkjhgfdfghjhgffghjkjhg"),
	})
	assert.NoError(t, err)

	err = manager.Create(tx, v0.User{
		PrimaryPublicKey:  []byte("dfghjkjhgrtyuiuytrfvbhy"),
		RecoveryPublicKey: []byte("fghjkjhgfdfghjhgffghjkjhg"),
	})
	assert.Error(t, err)
}

func TestManager_Alter(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	manager := newManager(t, db)

	tx, err := db.Begin()
	require.NoError(t, err)
	err = manager.Create(tx, v0.User{
		PrimaryPublicKey:  []byte("dfghjkjhgrtyuiuytrfvbhy"),
		RecoveryPublicKey: []byte("fghjkjhgfdfghjhgffghjkjhg"),
		Name:              "Laurens"})
	assert.NoError(t, err)

	err = manager.Alter(tx, v0.User{
		PrimaryPublicKey:  []byte("dfghjkjhgrtyuiuytrfvbhy"),
		RecoveryPublicKey: []byte("fghjkjhgfdfghjhgffghjkjhg"),
		Name:              "Karel"})
	assert.NoError(t, err)
}

func TestManager_Get(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	manager := newManager(t, db)

	tx, err := db.Begin()
	require.NoError(t, err)
	err = manager.Create(tx, v0.User{
		PrimaryPublicKey:  []byte("dfghjkjhgrtyuiuytrfvbhy"),
		RecoveryPublicKey: []byte("fghjkjhgfdfghjhgffghjkjhg"),
		Name:              "Laurens",
	})
	assert.NoError(t, err)

	user, err := manager.Get(tx, []byte("dfghjkjhgrtyuiuytrfvbhy"))
	require.NoError(t, err)

	assert.Equal(t, "Laurens", user.Name)
}
