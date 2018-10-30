package core

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type Router interface {
	Route(ctx context.Context, config *v0.Config) (*v0.Response, error)
	InfoService(ctx context.Context, e *empty.Empty) (info *v0.Info, err error)
}
