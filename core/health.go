package core

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"net/http"
	"time"
)

type HealthManager interface {
	GRPC(context.Context, *v0.HealthCheckRequest) (*v0.HealthCheckResponse, error)
	http.Handler
	Initial(check func() error)
	Async(check func() error, interval time.Duration)
	Start()
}
