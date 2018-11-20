package carapace

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

func (c *Client) FillUserKeys(config *v0.Config) error {
	switch config.Spec.TypeUrl {
	case OwnerSet:
		set := v0.OwnerSet{}
		err := ptypes.UnmarshalAny(config.Spec, &set)
		if err != nil {
			return err
		}
		err = c.fillOwnerSet(&set)
		if err != nil {
			return err
		}

		any, err := ptypes.MarshalAny(&set)
		if err != nil {
			return err
		}

		config.Spec = any
		return nil
	default:
		return errors.New("unrecognised Spec type URL")
	}
}

func (c *Client) fillOwnerSet(set *v0.OwnerSet) error {
	for _, owner := range set.Owners {
		if owner.Name == c.Name && owner.Email == c.Email {
			pubkey, err := c.PublicKey()
			if err != nil {
				return err
			}
			owner.PrimaryPublicKey = pubkey

			rpubkey, err := c.RPublicKey()
			if err != nil {
				return err
			}
			owner.RecoveryPublicKey = rpubkey
			return nil
		}
	}
	return errors.New("unable to find user in config file")
}
