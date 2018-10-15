package core

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type ConfigStore interface {
	Add(config v1.Config) (Response, error)
}
