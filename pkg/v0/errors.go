package v0_handler

import (
	"github.com/pkg/errors"
)

var (
	ErrIncorrectKind = errors.New("invalid kind provided to handler")
)
