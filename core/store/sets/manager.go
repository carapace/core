package sets

import (
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core/store/util"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var Sets = []proto.Message{
	&v0.OwnerSet{},
}

type Manager struct {
	OwnerSet OwnerSet
	UserSet  UserSet
}

var (
	ErrNotExist = errors.New("set not found")
)

func (m *Manager) AutoMigrate(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS owner_sets
			(
				ID INTEGER PRIMARY KEY, 
				created_at TIMESTAMP, 
				owner_set BLOB, 
				deleted_at TIMESTAMP NULL DEFAULT NULL 
			);`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`CREATE TABLE IF NOT EXISTS user_sets
			(
				ID INTEGER PRIMARY KEY, 
				created_at TIMESTAMP, 
				name string,
				user_set BLOB, 
				deleted_at TIMESTAMP NULL DEFAULT NULL 
			);`)
	if err != nil {
		return err
	}

	// This ensures we can soft delete models, but nothing else
	err = util.SetUpdateOnly(tx, "owner_sets")
	if err != nil {
		return err
	}

	return util.SetUpdateOnly(tx, "user_sets")
}
