package sets

import (
	"database/sql"
	"github.com/carapace/core/core/store/util"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up0000003, Down0000003)
}

func Up0000003(tx *sql.Tx) error {
	_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS identities
			(
				ID INTEGER PRIMARY KEY, 
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				name string, 
				identity BLOB, 
				deleted_at TIMESTAMP NULL DEFAULT NULL 
			);`)
	if err != nil {
		return err
	}

	// This ensures we can soft delete models, but nothing else
	err = util.SetUpdateOnly(tx, "identities")
	if err != nil {
		return err
	}

	return nil
}

func Down0000003(tx *sql.Tx) error {
	return nil
}
