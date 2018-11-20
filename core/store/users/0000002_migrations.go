package user

import (
	"database/sql"
	"github.com/carapace/core/core/store/util"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

var (
	ErrUserDoesNotExists = errors.New("user does not exist")
)

func init() {
	goose.AddMigration(Up0000002, Down0000002)
}

func Up0000002(tx *sql.Tx) error {
	_, err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS users (
				ID INTEGER PRIMARY KEY, 
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
				deleted_at TIMESTAMP NULL,
				deleted BOOLEAN DEFAULT FALSE,
				name string NOT NULL,
    			email string NOT NULL , 
    			primary_public_key BLOB NOT NULL,
  			   	recovery_public_key BLOB NOT NULL,
				
				user_set string,
				super_user BOOLEAN NOT NULL,
    			auth_level INT DEFAULT 0,
    			weight INT DEFAULT 0
    			);
    	  
				CREATE UNIQUE INDEX IF NOT EXISTS unique_user ON users(primary_public_key, deleted, recovery_public_key, name, email);
    	  `)
	if err != nil {
		return err
	}

	// This ensures we can soft delete models, but nothing else
	return util.SetUpdateOnly(tx, "users")
}

func Down0000002(tx *sql.Tx) error {
	return nil
}
