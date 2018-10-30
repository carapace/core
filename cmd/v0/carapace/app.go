package carapace

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/health"
	"github.com/carapace/core/pkg/router"
	"google.golang.org/grpc"
)

func New(loglevel string) (*core.App, *grpc.Server, error) {
	logger := defaultLogger(loglevel)
	defer logger.Sync()

	cfg, err := core.Config{
		Router:        router.Router,
		Logger:        logger,
		HealthManager: health.New(),

		Health: core.Health{
			Port: "9000",
			Host: "0.0.0.0",
		},
	}.Build()

	if err != nil {
		return nil, nil, err
	}

	app, err := core.New(cfg)
	if err != nil {
		return nil, nil, err
	}

	server := grpc.NewServer()
	v0.RegisterCoreServiceServer(server, app)
	return app, server, nil
}
