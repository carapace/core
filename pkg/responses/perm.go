package response

import (
	"github.com/carapace/core/api/v0/proto"
)

func PermissionDenied(context string) *v0.Response {
	return &v0.Response{Code: v0.Code_Forbidden, MSG: context}
}
