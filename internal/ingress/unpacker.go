package ingress

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type Unpacker interface {
	Unpack(config v1.Config)
}
