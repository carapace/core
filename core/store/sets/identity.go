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

const (
	IdentityURL = "type.googleapis.com/v0.Identity"
)

func (i *Identity) Put(ctx context.Context, tx *sql.Tx, set *v0.Identity) error {
	serialized, err := proto.Marshal(set)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
			UPDATE resources SET 
                     deleted_at = ?
			WHERE deleted_at = NULL AND name = ? AND proto_url = ?;`, time.Now(), set.Name, IdentityURL)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `
					INSERT INTO resources 
					  (created_at, resource, deleted_at, name, proto_url) 
					VALUES (?, ?, ?, ?, ?)`,
		time.Now(),
		serialized,
		nil,
		set.Name,
		IdentityURL,
	)
	return err
}

func (i *Identity) Get(ctx context.Context, tx *sql.Tx, name string) (*v0.Identity, error) {
	row := tx.QueryRowContext(ctx, `
			SELECT resource FROM resources 
			WHERE deleted_at IS NULL AND name = ? AND proto_url = ? ;`, name, IdentityURL)

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
	rows, err := tx.QueryContext(ctx, `SELECT resource FROM resources WHERE deleted_at IS NULL AND proto_url = ?;`, IdentityURL)
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
