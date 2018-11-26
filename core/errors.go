package core

import (
	"github.com/pkg/errors"
)

var (
	ErrUnAuthorized = errors.New("unauthorized action")
	ErrNotExists    = errors.New("not found")
	ErrNoOwners     = errors.New("no ownerSet found")
)
