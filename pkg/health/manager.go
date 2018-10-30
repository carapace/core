package health

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"net/http"
	"sync"
	"time"
)

var _ core.HealthManager = &Manager{}

type Manager struct {
	mu   *sync.RWMutex
	cron *cron.Cron

	healthy  bool
	initials []func() error
	ats      []func() error
}

func (m *Manager) Initial(check func() error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.initials = append(m.initials, check)
}

func (m *Manager) Async(check func() error, interval time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cron.Schedule(cron.Every(interval), cron.FuncJob(func() {
		err := check()
		if err != nil {
			m.healthy = false
			return
		}
		m.healthy = true
	}))
}

func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.healthy {
		http.Error(w, "", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (m *Manager) GRPC(context.Context, *v0.HealthCheckRequest) (*v0.HealthCheckResponse, error) {
	switch m.healthy {
	case false:
		return &v0.HealthCheckResponse{Status: v0.HealthCheckResponse_NOT_SERVING}, nil
	case true:
		return &v0.HealthCheckResponse{Status: v0.HealthCheckResponse_SERVING}, nil
	default:
		return nil, errors.New("health check unresponsive")
	}
}

func (m *Manager) Start() {
	for _, f := range m.initials {
		err := f()
		if err != nil {
			m.healthy = false
		}
	}
	m.cron.Start()
}

func New() *Manager {
	return &Manager{
		mu:       &sync.RWMutex{},
		cron:     cron.New(),
		healthy:  true,
		initials: []func() error{},
		ats:      []func() error{},
	}
}
