package util

import (
	"database/sql"
	"fmt"
)

// nolint: gosec
func SetUpdateOnly(tx *sql.Tx, tableName string) error {
	_, err := tx.Exec(fmt.Sprintf(`	
			CREATE TRIGGER IF NOT EXISTS %s_no_update 
			BEFORE UPDATE ON %s
			BEGIN
			  SELECT
    			CASE
 					WHEN NEW.deleted_at IS NULL THEN
 				RAISE ( ABORT, 'only updates on users are allowed')
              END;
			END;`, tableName, tableName))
	return err
}
