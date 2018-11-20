package response

import (
	"github.com/carapace/core/api/v0/proto"
)

func Err(err error) *v0.Response {
	return &v0.Response{Code: v0.Code_Internal, MSG: "", Err: err.Error()}
}
