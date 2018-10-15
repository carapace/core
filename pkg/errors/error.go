package errors

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type Error struct {
	Code v1.Code
	Err  string
}

func New(err string, code v1.Code) Error {
	return Error{Code: code, Err: err}
}

func (e Error) Error() string {
	return e.Err
}
