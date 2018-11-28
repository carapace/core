package policies

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
)

func (m *Manager) Add(ctx context.Context, tx *sql.Tx, policy *v0.Policy, resource, namespace string) error {

	serialized, err := json.Marshal(policy.Subjects)
	if err != nil {
		return err
	}

	affected, err := tx.ExecContext(ctx, `
		INSERT INTO policies (str_id, description, effect, actions, meta, deleted_at, subjects) VALUES (
		?, ?, ?, ?, ?, FALSE, ?); SELECT last_insert_rowid();
	`, policy.ID, policy.Description, policy.Effect.String(), SerializeActions(policy.Actions...), policy.Meta, serialized,
	)
	if err != nil {
		return err
	}

	policyID, err := affected.LastInsertId()
	if err != nil {
		return err
	}

	err = m.AddResource(ctx, tx, policyID, resource, namespace)
	if err != nil {
		return err
	}

	for _, condition := range policy.GetConditions() {
		_, err := m.CreateCondition(ctx, tx, condition, "GLOBAL", condition.Args)
		if err != nil {
			return errors.Wrap(err, "create condition")
		}

		err = m.AddCondition(ctx, tx, policyID, condition.String(), "GLOBAL")
		if err != nil {
			return errors.Wrap(err, "add condition PK")
		}
	}
	return err
}

func (m *Manager) AddSubject(ctx context.Context, tx *sql.Tx, policyID int64, username string) error {
	affected, err := tx.ExecContext(ctx, `INSERT INTO policies_users (policy_id, user_id) VALUES 
                                             (?, (SELECT ID FROM users WHERE name = ?))`, policyID, username)
	if err != nil {
		return err
	}
	rowsAffected, err := affected.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (m *Manager) AddResource(ctx context.Context, tx *sql.Tx, policyID int64, resource, namespace string) error {
	affected, err := tx.ExecContext(ctx, `INSERT INTO policies_resources (policy_id, resource_id) VALUES 
                                         (?, (SELECT ID FROM resources WHERE name = ? AND namespace = ?))`, policyID, resource, namespace)
	if err != nil {
		return err
	}
	rowsAffected, err := affected.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("resource not found")
	}

	return nil
}

func (m *Manager) CreateCondition(ctx context.Context, tx *sql.Tx, condition *v0.Condition, namespace string, constructor *any.Any) (int64, error) {
	serialed, err := proto.Marshal(constructor)
	if err != nil {
		return 0, err
	}

	affected, err := tx.ExecContext(ctx, `INSERT INTO conditions (name, constructor, namespace, kind) VALUES 
                                         (?, ?, ?, ?); SELECT last_insert_rowid()`, condition.String(), serialed, namespace, condition.Name)
	if err != nil {
		return 0, err
	}

	return affected.LastInsertId()
}

func (m *Manager) AddCondition(ctx context.Context, tx *sql.Tx, policyID int64, condition string, namespace string) error {
	affected, err := tx.ExecContext(ctx, `INSERT INTO policies_conditions (policy_id, condition_id) VALUES 
                                             (?, (SELECT ID from conditions WHERE name = ? AND namespace = ?))`, policyID, condition, namespace)
	if err != nil {
		return err
	}

	rowsAffected, err := affected.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("condition not found")
	}
	return err
}
