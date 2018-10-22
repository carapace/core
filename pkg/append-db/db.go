package append

import (
	"github.com/carapace/cellar"
	"github.com/carapace/core/pkg/state"
)

type DB struct {
	cellar *cellar.DB

	cache  Cache
	hasher state.Hasher
	signer state.Signer
}

func New(folder string, opts ...Option) (*DB, error) {
	c, err := cellar.New(folder)
	if err != nil {
		return nil, err
	}

	db := &DB{
		cellar: c,
	}

	for _, opt := range opts {
		err := opt(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func (db *DB) Close() error {
	return db.cellar.Close()
}
