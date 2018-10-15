package core

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/pkg/errors"
	"github.com/carapace/core/pkg/sanitizer"
)

// ResponseSuccess returns a formatted response indicating a successful API call.
// MSG should be a meaningful message, not "everything went fine",
// since that is indicated by the Code field and lack of Err field.
func ResponseSuccess(MSG string) *v1.Response {
	return &v1.Response{
		MSG:  MSG,
		Code: v1.Code_OK,
		Err:  "",
	}
}

// ResponseErrInternal returns a formatted response indicating a failed API call
// due to an internal error.
func ResponseErrInternal(err string) *v1.Response {
	return &v1.Response{
		MSG:  "",
		Code: v1.Code_Internal,
		Err:  sanitize.String(err),
	}
}

// ResponseFromErr returns a formatted response indicating a failed API call
// due to an error. The error message is sanitized before being returned
func ResponseFromErr(err error) *v1.Response {
	switch err.(type) {
	case errors.Error:
		e := err.(errors.Error)
		return &v1.Response{
			Code: e.Code,
			Err:  sanitize.Error(e).Error(),
			MSG:  "",
		}
	default:
		return ResponseErrInternal(sanitize.Error(err).Error())
	}
}
