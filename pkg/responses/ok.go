package response

import (
	"github.com/carapace/core/api/v0/proto"
)

func OK(msg string) *v0.Response {
	return &v0.Response{MSG: msg, Code: v0.Code_OK, Err: ""}
}
