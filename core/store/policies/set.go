package policies

import (
	"context"
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"time"
)

func (m *Manager) Set(ctx context.Context, tx *sql.Tx, policies []*v0.Policy, resource, namespace string) error {
	_, err := tx.ExecContext(ctx, `
						UPDATE policies SET deleted_at = ? 
						WHERE ID = (SELECT policy_id FROM policies_resources 
							WHERE resource_id = (SELECT ID FROM resources WHERE name = ? AND namespace = ?))`,
		time.Now(), resource, namespace)
	if err != nil {
		return err
	}

	for _, policy := range policies {
		err = m.Add(ctx, tx, policy, resource, namespace)
		if err != nil {
			return err
		}
	}
	return nil
}
