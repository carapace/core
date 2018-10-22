package append

import (
	"github.com/carapace/cellar"
	"github.com/carapace/core/pkg/state"
)

type Option func(db *DB) error

func WithCache(cache Cache) Option {
	return func(db *DB) error {
		db.cache = cache
		return nil
	}
}

func WithHasher(hasher state.Hasher) Option {
	return func(db *DB) error {
		db.hasher = hasher
		return nil
	}
}

func WithSigner(signer state.Signer) Option {
	return func(db *DB) error {
		db.signer = signer
		return nil
	}
}

func WithCipher(cipher cellar.Cipher) Option {
	return func(db *DB) error {
		return cellar.WithCipher(cipher)(db.cellar)
	}
}
