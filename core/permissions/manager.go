package permissions

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
)

type Manager struct {
	store *core.Store
	ladon ladon.Ladon
	core.PolicyManager
}

func New(store *core.Store) *Manager {
	return &Manager{store: store, ladon: ladon.Ladon{Manager: memory.NewMemoryManager()}, PolicyManager: store.Policies}
}

func (m *Manager) IsAllowed(ctx context.Context, tx *sql.Tx, namespace, resource string, request *ladon.Request) error {
	rawpolices, err := m.store.Policies.Get(ctx, tx, resource, namespace)

	if err != nil {
		return err
	}

	return m.PoliciesAllow(request, rawpolices, namespace)
}

func (m *Manager) PoliciesAllow(request *ladon.Request, rawpolicies []*v0.Policy, namespace string) error {
	policies, err := PoliciesFromProto(rawpolicies, fmt.Sprintf("%s.%s", namespace, request.Resource))
	if err != nil {
		return err
	}
	request.Resource = fmt.Sprintf("%s.%s", namespace, request.Resource)
	return m.ladon.DoPoliciesAllow(request, policies)
}
