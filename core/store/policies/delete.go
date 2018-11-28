package policies

import (
	"context"
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"github.com/pkg/errors"
	"time"
)

func (m *Manager) Delete(ctx context.Context, tx *sql.Tx, policy *v0.Policy) error {
	affected, err := tx.ExecContext(ctx, `
		UPDATE policies SET deleted_at = ? WHERE str_id = ?;`, time.Now(), policy.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := affected.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 0 {
		return errors.New("policy not found")
	}
	return nil
}
