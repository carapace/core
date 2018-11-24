package sets

import (
	"context"
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"time"
)

type Identity struct{}

func (i *Identity) Put(ctx context.Context, tx *sql.Tx, set *v0.Identity) error {
	serialized, err := proto.Marshal(set)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE identities SET deleted_at = ? WHERE deleted_at = NULL AND name = ?;`, time.Now(), set.Name)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO identities (created_at, identity, deleted_at, name) VALUES (?, ?, ?, ?)`, time.Now(), serialized, nil, set.Name)
	return err
}

func (i *Identity) Get(ctx context.Context, tx *sql.Tx, name string) (*v0.Identity, error) {
	row := tx.QueryRowContext(ctx, `SELECT identity FROM identities WHERE deleted_at IS NULL AND name = ?;`, name)

	var data = []byte{}
	err := row.Scan(&data)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			return nil, errors.New("no matching identity found")
		}
		return nil, err
	}

	var set = &v0.Identity{}
	err = proto.Unmarshal(data, set)
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (i *Identity) All(ctx context.Context, tx *sql.Tx) ([]*v0.Identity, error) {
	rows, err := tx.QueryContext(ctx, `SELECT identity FROM identities WHERE deleted_at IS NULL;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []*v0.Identity{}
	for rows.Next() {
		set := v0.Identity{}
		data := []byte{}
		err := rows.Scan(&data)
		if err != nil {
			return nil, err
		}

		err = proto.Unmarshal(data, &set)
		if err != nil {
			return nil, err
		}
		res = append(res, &set)
	}
	return res, nil
}
