//go:generate mockgen -destination=mocks/warden_mock.go -package=mock github.com/carapace/core/core Warden
//go:generate mockgen -destination=mocks/policymanager_mock.go -package=mock github.com/carapace/core/core PolicyManager
//go:generate mockgen -destination=mocks/permissionmanager_mock.go -package=mock github.com/carapace/core/core PermissionManager

package core

import (
	"context"
	"database/sql"

	"github.com/carapace/core/api/v0/proto"
	"github.com/ory/ladon"
)

type PermissionManager interface {
	PolicyManager
	Warden
}

type Warden interface {
	IsAllowed(ctx context.Context, tx *sql.Tx, resource, namespace string, request *ladon.Request) error
	PoliciesAllow(request *ladon.Request, rawpolicies []*v0.Policy, namespace string) error
}

type PolicyManager interface {
	Add(ctx context.Context, tx *sql.Tx, policy *v0.Policy, resource, namespace string) error
	Set(ctx context.Context, tx *sql.Tx, policies []*v0.Policy, resource, namespace string) error
	Get(ctx context.Context, tx *sql.Tx, resource, namespace string) ([]*v0.Policy, error)
	Delete(ctx context.Context, tx *sql.Tx, policy *v0.Policy) error
	All(ctx context.Context, tx *sql.Tx) ([]*v0.Policy, error)
}
