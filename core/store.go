//go:generate mockgen -destination=mocks/userstore_mock.go -package=mock github.com/carapace/core/core UserStore
//go:generate mockgen -destination=mocks/ownerset_mock.go -package=mock github.com/carapace/core/core OwnerSet
//go:generate mockgen -destination=mocks/userset_mock.go -package=mock github.com/carapace/core/core UserSet
//go:generate mockgen -destination=mocks/configmanager_mock.go -package=mock github.com/carapace/core/core ConfigManager
//go:generate mockgen -destination=mocks/identityset_mock.go -package=mock github.com/carapace/core/core IdentitySet

package core

import (
	"context"
	"database/sql"
	"fmt"
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

	fmt.Println("STORE INIT", manager.Sets.Config)

	return &Store{
		DB: db,
		Sets: &Sets{
			OwnerSet: manager.Sets.OwnerSet,
			UserSet:  manager.Sets.UserSet,
			Config:   manager.Sets.Config,
			Identity: manager.Sets.Identity,
		},
		Users: manager.Users,
	}, nil
}

type Store struct {
	DB    *sql.DB
	Sets  *Sets
	Users UserStore
}

func (s *Store) Begin() (*sql.Tx, error) {
	return s.DB.Begin()
}

type Sets struct {
	OwnerSet OwnerSet
	UserSet  UserSet
	Config   ConfigManager
	Identity IdentitySet
}

type OwnerSet interface {
	Get(context.Context, *sql.Tx) (*v0.OwnerSet, error)
	Put(context.Context, *sql.Tx, *v0.OwnerSet) error
}

type UserSet interface {
	Get(context.Context, *sql.Tx, string) (*v0.UserSet, error)
	Put(context.Context, *sql.Tx, *v0.UserSet) error
	All(context.Context, *sql.Tx) ([]*v0.UserSet, error)
}

type UserStore interface {
	Create(context.Context, *sql.Tx, v0.User) error
	Alter(context.Context, *sql.Tx, v0.User) error
	Get(ctx context.Context, tx *sql.Tx, publicKey []byte) (*v0.User, error)
	Delete(context.Context, *sql.Tx, v0.User) error
	BySet(context.Context, *sql.Tx, string) ([]*v0.User, error)
}

type ConfigManager interface {
	Add(context.Context, *sql.Tx, *v0.Config) error
}

type IdentitySet interface {
	Get(context.Context, *sql.Tx, string) (*v0.Identity, error)
	Put(context.Context, *sql.Tx, *v0.Identity) error
	All(context.Context, *sql.Tx) ([]*v0.Identity, error)
}
