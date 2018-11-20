//go:generate mockgen -destination=mocks/appendstore_mock.go -package=mock github.com/carapace/core/core SetStore
//go:generate mockgen -destination=mocks/userstore_mock.go -package=mock github.com/carapace/core/core UserStore

package core

import (
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core/store"
	"github.com/pkg/errors"
)

var (
	ErrNoOwners = errors.New("no ownerSet found")
)

type StoreAPI struct {
	*store.Manager
}

func NewStore(db *sql.DB) (*StoreAPI, error) {
	m, err := store.New(db)
	if err != nil {
		return nil, err
	}
	return &StoreAPI{m}, nil
}

type SetStore interface {
	GetOwnerSet(*sql.Tx) (*v0.OwnerSet, error)
	PutOwnerSet(tx *sql.Tx, set *v0.OwnerSet) error
}

type UserStore interface {
	Create(tx *sql.Tx, user v0.User) error
	Alter(tx *sql.Tx, user v0.User) error
	Get(tx *sql.Tx, publicKey string) (*v0.User, error)
	Delete(tx *sql.Tx, publicKey string) error
	AlterOrCreate(TX, user v0.User) error
}
