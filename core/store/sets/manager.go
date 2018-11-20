package sets

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var Sets = []proto.Message{
	&v0.OwnerSet{},
}

type Manager struct {
	OwnerSet *OwnerSet
	UserSet  *UserSet
}

var (
	ErrNotExist = errors.New("set not found")
)
