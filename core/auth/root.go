package auth

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (m *Manager) GrantRoot(witness *v0.Witness) (bool, error) {

	tx, err := m.Store.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	weight, err := m.Quorum()
	if err != nil {
		return false, err
	}

	var totalWeight int32
	for _, sig := range witness.Signatures {
		user, err := m.Store.Users.Get(tx, sig.GetPrimaryPublicKey())
		if err != nil {
			core.Logger.Info("user does not exist")
			// user does not exist (so a new user is being created)
			if err.Error() == "sql: no rows in result set" {
				continue
			}
			return false, err
		}

		core.Logger.Info("", zap.String("USER", user.GetName()))

		totalWeight += user.Weight
	}
	core.Logger.Info("total weight: ", zap.Int32("WEIGHT", weight), zap.Int32("TOTAL WEIGHT", totalWeight))
	return totalWeight >= weight, nil
}

func (m *Manager) GrantBackupRoot(witness *v0.Witness) (bool, error) {
	tx, err := m.Store.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	weight, err := m.Quorum()
	if err != nil {
		return false, err
	}

	var totalWeight int32
	for _, sig := range witness.Signatures {
		var err error
		var user *v0.User

		if sig.GetRecoveryPublicKey() != nil {
			user, err = m.Store.Users.Get(tx, sig.GetRecoveryPublicKey())
		} else {
			user, err = m.Store.Users.Get(tx, sig.GetPrimaryPublicKey())
		}

		if err != nil {
			// user does not exist (so a new user is being created)
			if err.Error() == "sql: no rows in result set" {
				continue
			}
			return false, err
		}

		totalWeight += user.Weight
	}

	return totalWeight >= weight, nil
}

func (m *Manager) SetOwners(set *v0.OwnerSet) error {
	tx, err := m.Store.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = m.Store.Sets.OwnerSet.Put(tx, set)
	if err != nil {
		return err
	}
	return errors.Wrapf(tx.Commit(), "auth.SetOwners")
}
