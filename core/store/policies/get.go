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

func (m *Manager) Get(ctx context.Context, tx *sql.Tx, resource, namespace string) (res []*v0.Policy, err error) {
	policies, err := tx.QueryContext(ctx, `
				SELECT ID, str_id, description, effect, actions, meta, subjects FROM policies 
				WHERE ID IN (SELECT policy_id FROM policies_resources 
					WHERE (SELECT ID FROM resources WHERE name = ? AND namespace = ?)) AND deleted_at = FALSE`, resource, namespace)
	if err != nil {
		return nil, err
	}
	defer policies.Close()

	for policies.Next() {
		var policyID int64 = 0
		var policy = &v0.Policy{}
		var effect string
		var actions string
		var subjects []byte
		err = policies.Scan(&policyID, &policy.ID, &policy.Description, &effect, &actions, &policy.Meta, &subjects)
		if err != nil {
			return nil, errors.Wrap(err, "policies scanning")
		}

		policy.Effect = v0.Effect(v0.Effect_value[effect])
		policy.Actions = DeserializeActions(actions)
		err := json.Unmarshal(subjects, &policy.Subjects)

		conditions, err := tx.QueryContext(ctx,
			`SELECT kind, constructor FROM conditions 
				   WHERE ID IN (SELECT condition_id FROM policies_conditions WHERE policy_id = ?)`, policyID)
		if err != nil {
			return nil, err
		}
		defer conditions.Close()

		for conditions.Next() {
			var cond = &v0.Condition{Args: &any.Any{}}
			var serializedArgs = []byte{}
			err = conditions.Scan(&cond.Name, &serializedArgs)
			if err != nil {
				return nil, err
			}

			err := proto.Unmarshal(serializedArgs, cond.Args)
			if err != nil {
				return nil, err
			}
			policy.Conditions = append(policy.Conditions, cond)
		}
		res = append(res, policy)
	}
	return
}
