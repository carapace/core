package core

import (
	"context"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (a *App) ConfigService(ctx context.Context, config *v0.Config) (res *v0.Response, err error) {
	defer func() {
		if r := recover(); r != nil {
			res = &v0.Response{Code: v0.Code_Internal, MSG: "", Err: fmt.Sprintf("%v", r)}
			err = nil
		}
	}()

	a.Logger.Info("Received new configuration object",
		zap.String("apiVersion", config.Header.ApiVersion),
		zap.String("kind", config.Header.Kind),
		zap.String("config", config.String()),
	)

	res, err = a.Router.Route(ctx, config)
	if err != nil {
		a.Logger.Warn("Error while handling new configuration object",
			zap.Error(err),
			zap.String("message", res.MSG),
			zap.Uint32("code", uint32(res.Code)),
		)

		panic(err) // caught by the top level defer
	}
	return res, err
}

func (a *App) InfoService(ctx context.Context, e *empty.Empty) (info *v0.Info, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("internal server error")
		}
	}()

	a.Logger.Info("Received request for node info")
	return a.Router.InfoService(ctx, e)
}

func (a *App) Check(ctx context.Context, req *v0.HealthCheckRequest) (res *v0.HealthCheckResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("internal server error")
		}
	}()

	a.Logger.Debug("Received request health check")
	return a.HealthManager.GRPC(ctx, req)
}
