package sets

import (
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
	"time"
)

type OwnerSet struct{}

func (o *OwnerSet) Put(tx *sql.Tx, set *v0.OwnerSet) error {
	serialized, err := proto.Marshal(set)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE owner_sets SET deleted_at = ? WHERE deleted_at = NULL`, time.Now())
	if err != nil {
		return err
	}
	_, err = tx.Exec(`INSERT INTO owner_sets (created_at, owner_set, deleted_at) VALUES (?, ?, ?)`, time.Now(), serialized, nil)
	return err
}

func (o *OwnerSet) Get(tx *sql.Tx) (*v0.OwnerSet, error) {
	row := tx.QueryRow(`SELECT owner_set FROM owner_sets WHERE deleted_at IS NULL;`)

	var data = []byte{}
	err := row.Scan(&data)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			return nil, ErrNotExist
		}
		return nil, err
	}

	var set = &v0.OwnerSet{}
	err = proto.Unmarshal(data, set)
	if err != nil {
		return nil, err
	}
	return set, nil
}
