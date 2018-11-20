package chaindb

import (
	"fmt"
	"sync"
	"time"

	"github.com/boltdb/bolt"
	"github.com/carapace/cellar"
	"github.com/carapace/core/pkg/state"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// DB is the god level struct holding chain-db
type DB struct {
	config Config
	mu     *sync.RWMutex
}

// Config contains user provided settings. Calling build will validate the config and fill defaults,
// or return an error if the user must provide the setting.
type Config struct {
	Folder string
	Logger *zap.Logger
	MetaDB cellar.MetaDB

	Hasher   state.Hasher
	Signer   state.Signer
	Verifier state.Verifier
	Cache    Cache
	Store    StorageEngine
}

// Build validates the configuration and sets defaults if not provided.
func (c Config) Build() error {
	if c.Folder == "" {
		return errors.New("chain-db: invalid folder")
	}

	if c.Logger == nil {
		c.Logger = defaultLogger()
	}

	if c.MetaDB == nil {
		metadb, err := bolt.Open(
			fmt.Sprintf("%s/%s", c.Folder, "meta.bolt"),
			0600,
			&bolt.Options{Timeout: 2 * time.Second},
		)
		if err != nil {
			return err
		}
		c.MetaDB = &cellar.BoltMetaDB{DB: metadb}
		err = c.MetaDB.Init()
		if err != nil {
			return err
		}
	}

	if c.Hasher == nil {
		return errors.New("chain-db: no hasher provided")
	}

	if c.Signer == nil {
		return errors.New("chain-db: no signer provided")
	}

	if c.Cache == nil {
		return errors.New("chain-db: no cache provided")
	}

	if c.Verifier == nil {
		return errors.New("chain-db: no signature validator provided")
	}

	if c.Store == nil {
		// init sqlite storage
	}

	return nil
}

// New is the constructor for config
func New(config Config, options ...ConfOption) (*DB, error) {
	db := &DB{
		config: config,
		mu:     &sync.RWMutex{},
	}

	var err error
	for _, opt := range options {
		err = opt(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func (db *DB) Close() error {
	return db.config.Store.Close()
}

func (db *DB) Logger() *zap.Logger {
	return db.config.Logger
}
