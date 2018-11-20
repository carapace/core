package core

import (
	"context"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service interface {
	Serve() error
}

func (a *App) ConfigService(ctx context.Context, config *v0.Config) (res *v0.Response, err error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		res = &v0.Response{Code: v0.Code_Internal, MSG: "", Err: fmt.Sprintf("%v", r)}
	// 		err = nil
	// 	}
	// }()

	Logger.Info("Received new configuration object",
		zap.String("apiVersion", config.Header.ApiVersion),
		zap.String("kind", config.Header.Kind),
		zap.String("config", config.String()),
	)

	res, err = a.Router.Route(ctx, config)
	if err != nil {
		res = &v0.Response{
			MSG:  "internal error: " + err.Error(),
			Code: v0.Code_Internal,
		}
	}
	return res, nil
}

func (a *App) InfoService(ctx context.Context, in *empty.Empty) (info *v0.RepeatedInfo, err error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		err = r.(error)
	// 	}
	// }()

	Logger.Info("Received request for node info")
	return a.Router.InfoService(ctx, in)
}

func (a *App) Check(ctx context.Context, req *v0.HealthCheckRequest) (res *v0.HealthCheckResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("internal server error")
		}
	}()

	Logger.Debug("Received request health check")
	return a.HealthManager.GRPC(ctx, req)
}
