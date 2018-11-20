//go:generate mockgen -destination=mocks/api_service_mock.go -package=mock github.com/carapace/core/core APIService

package core

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/carapace/core/api/v0/proto"
)

type Router interface {
	Route(ctx context.Context, config *v0.Config) (*v0.Response, error)
	Register(handlers ...APIService)
	InfoService(ctx context.Context, e *empty.Empty) (info *v0.RepeatedInfo, err error)
}

type APIService interface {
	ConfigService(context.Context, *v0.Config) (*v0.Response, error)
	InfoService() (*v0.Info, error)
}
