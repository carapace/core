package sets

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
)

type Config struct{}

func (c Config) Add(ctx context.Context, tx *sql.Tx, config *v0.Config) error {
	serialized, err := proto.Marshal(config)
	if err != nil {
		return err
	}

	hash := sha256.New()
	_, err = hash.Write(serialized)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
					INSERT INTO config_sets (incrementID, hash, config_set) VALUES 
					(?, ?, ?)`, config.Header.Increment, hash.Sum(nil), serialized)
	return err
}
