package condition

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
)

// InSets validates that the user is within a set
type Multisig struct {
	keys [][]byte
}

// Fulfills returns true the provided user has an AuthLevel >= the conditions level
func (c Multisig) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(*v0.Witness)

	set := make(map[string]struct{})

	for _, n := range s.Signatures {
		set[string(n.GetPrimaryPublicKey())] = struct{}{}
	}

	for _, k := range c.keys {
		if _, in := set[string(k)]; !in {
			return false
		}
	}
	return ok
}

// GetName returns the UserAuthGreater condition's name.
func (c Multisig) GetName() string {
	return "Multisig"
}

func newMultisig(any *any.Any) (ladon.Condition, error) {
	arg := &v0.MultiSigArg{}
	err := ptypes.UnmarshalAny(any, arg)
	if err != nil {
		return nil, err
	}
	return Multisig{keys: arg.PrimaryPublicKeys}, nil
}
