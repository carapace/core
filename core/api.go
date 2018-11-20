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
	a.mu.Lock()
	defer a.mu.Unlock()

	Logger.Info("Received new configuration object",
		zap.String("apiVersion", config.Header.ApiVersion),
		zap.String("kind", config.Header.Kind),
		zap.String("config", config.String()),
	)

	tx, err := a.Store.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	err = a.Store.Sets.Config.Add(tx, config)
	if err != nil {
		panic(err) // TODO create proper response specifying why the config is incorrect (has to do with incrementID)
	}

	res, err = a.Router.Route(ctx, config, tx)
	if err != nil {
		res = &v0.Response{
			MSG:  "internal error: " + err.Error(),
			Code: v0.Code_Internal,
		}
	}

	err = tx.Commit()
	if err != nil {
		a.Logger.Error("error while committing config transaction", zap.Error(err))
		panic(err)
	}
	return res, nil
}

func (a *App) InfoService(ctx context.Context, in *empty.Empty) (info *v0.RepeatedInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	a.mu.RLock()
	defer a.mu.RUnlock()

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
