package append

import (
	"github.com/carapace/cellar"
)

type ConfOption func(*DB) error

func WithCellarOption(option cellar.Option) ConfOption {
	return func(db *DB) error {
		return option(db.cellar)
	}
}
