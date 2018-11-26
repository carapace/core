package core

import (
	"context"
	"net/http"
	"time"

	"github.com/carapace/core/api/v0/proto"
)

type HealthManager interface {
	GRPC(context.Context, *v0.HealthCheckRequest) (*v0.HealthCheckResponse, error)
	http.Handler
	Initial(check func() error)
	Async(check func() error, interval time.Duration)
	Start()
}
