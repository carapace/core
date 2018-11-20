package sets

import (
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"time"
)

type UserSet struct{}

func (u *UserSet) Put(tx *sql.Tx, set *v0.UserSet) error {
	serialized, err := proto.Marshal(set)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE user_sets SET deleted_at = ? WHERE deleted_at = NULL AND name = ?;`, time.Now(), set.Set)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`INSERT INTO user_sets (created_at, user_set, deleted_at, name) VALUES (?, ?, ?, ?)`, time.Now(), serialized, nil, set.Set)
	return err
}

func (u *UserSet) Get(tx *sql.Tx, name string) (*v0.UserSet, error) {
	row := tx.QueryRow(`SELECT user_set FROM user_sets WHERE deleted_at IS NULL AND name = ?;`, name)

	var data = []byte{}
	err := row.Scan(&data)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			return nil, errors.New("no userSet found")
		}
		return nil, err
	}

	var set = &v0.UserSet{}
	err = proto.Unmarshal(data, set)
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (u *UserSet) All(tx *sql.Tx) ([]*v0.UserSet, error) {
	rows, err := tx.Query(`SELECT user_set FROM user_sets WHERE deleted_at IS NULL;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []*v0.UserSet{}
	for rows.Next() {
		set := v0.UserSet{}
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
