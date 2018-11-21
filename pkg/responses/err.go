package response

import (
	"fmt"
	"github.com/carapace/core/api/v0/proto"
)

func Err(err error) *v0.Response {
	return &v0.Response{Code: v0.Code_Internal, MSG: "", Err: err.Error()}
}

func ValidationErr(err error) *v0.Response {
	return &v0.Response{
		Code: v0.Code_BadRequest,
		MSG:  "validation rules may be found in the protocol definitions",
		Err:  fmt.Sprintf("proto validation: %s", err.Error()),
	}
}
