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

func NewStore(db *sql.DB) (*Store, error) {
	manager := store.New(db)

	err := manager.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
		Sets: Sets{
			OwnerSet: manager.Sets.OwnerSet,
			UserSet:  manager.Sets.UserSet,
		},
		Users: manager.Users,
	}, nil
}

type Store struct {
	DB *sql.DB
	Sets
	Users UserStore
}

func (s *Store) Begin() (*sql.Tx, error) {
	return s.DB.Begin()
}

type Sets struct {
	OwnerSet OwnerSet
	UserSet  UserSet
}

type OwnerSet interface {
	Get(*sql.Tx) (*v0.OwnerSet, error)
	Put(tx *sql.Tx, set *v0.OwnerSet) error
}

type UserSet interface {
	Get(*sql.Tx, string) (*v0.UserSet, error)
	Put(tx *sql.Tx, set *v0.UserSet) error
	All(*sql.Tx) ([]*v0.UserSet, error)
}

type UserStore interface {
	Create(tx *sql.Tx, user v0.User) error
	Alter(tx *sql.Tx, user v0.User) error
	Get(tx *sql.Tx, publicKey []byte) (*v0.User, error)
	Delete(tx *sql.Tx, user v0.User) error
	BySet(tx *sql.Tx, set string) ([]*v0.User, error)
}
