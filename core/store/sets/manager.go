package sets

import (
	"github.com/pkg/errors"
)

type Manager struct {
	OwnerSet *OwnerSet
	UserSet  *UserSet
	Config   *Config
}

func New() *Manager {
	return &Manager{
		OwnerSet: &OwnerSet{},
		UserSet:  &UserSet{},
		Config:   &Config{},
	}
}

var (
	ErrNotExist = errors.New("set not found")
)
