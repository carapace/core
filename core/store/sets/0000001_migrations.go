package sets

import (
	"database/sql"
	"github.com/carapace/core/core/store/util"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up0000001, Down0000001)
}

func Up0000001(tx *sql.Tx) error {
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

func Down0000001(tx *sql.Tx) error {
	return nil
}
