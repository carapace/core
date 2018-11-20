package response

import (
	"github.com/carapace/core/api/v0/proto"
)

func MSG(code v0.Code, msg string) *v0.Response {
	return &v0.Response{Code: code, MSG: msg}
}
