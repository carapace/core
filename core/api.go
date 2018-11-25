package core

import (
	"context"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/responses"
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

	// Validation for the config protocol buffer is autogenerated. See the source file for the validation rules.
	// Handlers may impose further validation.
	err = config.Validate()
	if err != nil {
		return response.ValidationErr(err), nil
	}

	Logger.Info("Received new configuration object",
		zap.String("apiVersion", config.Header.ApiVersion),
		zap.String("kind", config.Header.Kind),
		zap.String("config", config.String()),
	)

	tx, err := a.Store.Begin(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback() // rollback silently fails, unless a handler forgets to rollback or commit.

	err = a.Store.Sets.Config.Add(ctx, tx, config)
	if err != nil {
		panic(err) // TODO create proper response specifying why the config is incorrect (has to do with incrementID)
	}

	// add the TX to the running context
	ctx = ContextWithTransaction(ctx, tx)

	// the handler is responsible for closing committing the transaction
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
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("internal server error")
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

func (a *App) TransactionService(ctx context.Context, trs *v0.Transaction) (res *v0.Transaction, err error) {
	defer func() {
		if r := recover(); r != nil {
			res.Response = &v0.Response{Code: v0.Code_Internal, MSG: "", Err: fmt.Sprintf("%v", r)}
			err = nil
		}
	}()
	a.mu.RLock()
	defer a.mu.RUnlock()

	err = trs.Validate()
	if err != nil {
		res.Response = response.ValidationErr(err)
		return res, nil
	}

	tx, err := a.Store.Begin(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback() // rollback silently fails, unless a handler forgets to rollback or commit.

	// add the TX to the running context
	ctx = ContextWithTransaction(ctx, tx)

	return a.TXService.TransactionService(ctx, trs)
}
