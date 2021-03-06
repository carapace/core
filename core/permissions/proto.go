package permissions

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core/permissions/conditions"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
)

func PolicyFromProto(perm *v0.Policy, resources ...string) (*ladon.DefaultPolicy, error) {
	var actions []string
	for _, action := range perm.Actions {
		actions = append(actions, action.String())
	}

	var conditions = make(map[string]ladon.Condition)
	for _, cond := range perm.Conditions {

		factory, have := condition.Factory[cond.Name.String()]
		if !have {
			return nil, errors.New("condition not recognised")
		}

		var err error
		conditions[cond.Name.String()], err = factory(cond.Args)
		if err != nil {
			return nil, err
		}
	}

	return &ladon.DefaultPolicy{
		Description: perm.Description,
		Effect:      condition.Effects[perm.Effect],
		Subjects:    perm.Subjects,
		Actions:     actions,
		Conditions:  conditions,
		Resources:   resources,
	}, nil
}

func PoliciesFromProto(perms []*v0.Policy, resources ...string) (ladon.Policies, error) {
	var res ladon.Policies
	for _, perm := range perms {
		pol, err := PolicyFromProto(perm, resources...)
		if err != nil {
			return nil, err
		}
		res = append(res, pol)
	}
	return res, nil
}
