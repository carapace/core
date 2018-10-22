package sanitize

import (
	"github.com/pkg/errors"
)

func String(s string) string {
	return s
}

func Error(err error) error {
	return errors.New(String(err.Error()))
}
