package ownerset

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// unmarshalAny unmarshals the embedded any in a v0.Config and returns the ownerset
func unmarshalAny(any *any.Any) (*v0.OwnerSet, error) {
	var ownerset = &v0.OwnerSet{}
	err := ptypes.UnmarshalAny(any, ownerset)
	return ownerset, err
}
