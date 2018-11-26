package sets

import (
	"context"
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
	"time"
)

type OwnerSet struct{}

const (
	OwnerSetURL = "type.googleapis.com/v0.OwnerSet"
)

func (o *OwnerSet) Put(ctx context.Context, tx *sql.Tx, set *v0.OwnerSet) error {
	serialized, err := proto.Marshal(set)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE resources SET deleted_at = ? WHERE deleted_at = NULL AND proto_url = ?`, time.Now(), OwnerSetURL)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO resources (created_at, resource, deleted_at, proto_url) VALUES (?, ?, ?, ?)`, time.Now(), serialized, nil, OwnerSetURL)
	return err
}

func (o *OwnerSet) Get(ctx context.Context, tx *sql.Tx) (*v0.OwnerSet, error) {
	row := tx.QueryRowContext(ctx, `SELECT resource FROM resources WHERE deleted_at IS NULL AND proto_url = ?;`, OwnerSetURL)

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
