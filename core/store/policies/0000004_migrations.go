package policies

import (
	"database/sql"
	"github.com/carapace/core/core/store/util"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up0000004, Down0000004)
}

func Up0000004(tx *sql.Tx) error {
	_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS policies
			(
				ID INTEGER PRIMARY KEY,
				str_id string,
				description string,
				effect string,
				actions string,
				meta BLOB,
				subjects string,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
				deleted_at TIMESTAMP NULL DEFAULT NULL ,
			
-- 				FOREIGN KEY(ID) REFERENCES policies_users(policy_id),
				FOREIGN KEY(ID) REFERENCES policies_resources(policy_id),
				FOREIGN KEY(ID) REFERENCES policies_conditions(policy_id)
			);

			CREATE TABLE IF NOT EXISTS conditions (
			  	ID INTEGER PRIMARY KEY,
			  	namespace string DEFAULT 'GLOBAL', 
			  	name string,
			  	kind INTEGER,
			  	constructor BLOB
			); 

-- 			CREATE TABLE IF NOT EXISTS policies_users (
--                	policy_id INTEGER NOT NULL, 
--                	user_id INTEGER NOT NULL,
--                 FOREIGN KEY(policy_id) REFERENCES policies(ID),
--                 FOREIGN KEY(user_id) REFERENCES users(ID)
-- 			); 

			CREATE TABLE IF NOT EXISTS policies_conditions (
               	policy_id INTEGER NOT NULL , 
               	condition_id INTEGER NOT NULL ,
                FOREIGN KEY(policy_id) REFERENCES policies(ID),
                FOREIGN KEY(condition_id) REFERENCES conditions(ID)
			);

			CREATE TABLE IF NOT EXISTS policies_resources (
               	policy_id INTEGER NOT NULL ,
               	resource_id INTEGER NOT NULL,
                FOREIGN KEY(policy_id) REFERENCES policies(ID),
                FOREIGN KEY(resource_id) REFERENCES resources(ID)
			); 
`)
	if err != nil {
		return err
	}

	// This ensures we can soft delete models, but nothing else
	err = util.SetUpdateOnly(tx, "policies")
	if err != nil {
		return err
	}
	return nil
}

func Down0000004(tx *sql.Tx) error {
	return nil
}
