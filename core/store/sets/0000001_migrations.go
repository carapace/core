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
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
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
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
				name string,
				user_set BLOB, 
				deleted_at TIMESTAMP NULL DEFAULT NULL 
			);`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`CREATE TABLE IF NOT EXISTS config_sets
			(
				ID INTEGER PRIMARY KEY, 
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
				incrementID INT,
				hash BLOB,
				config_set BLOB
			);
			
			CREATE UNIQUE INDEX IF NOT EXISTS config_sets_unique ON config_sets(hash, incrementID);
			`)
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
