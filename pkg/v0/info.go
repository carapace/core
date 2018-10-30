package v0_handler

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type Info struct {
	Info *v0.Info
}

func (i Info) InfoService(ctx context.Context, e *empty.Empty) (info *v0.Info, err error) {
	return i.Info, nil
}
