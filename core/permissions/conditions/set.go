package condition

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
)

// InSets validates that the user is within a set
type InSets struct {
	sets []string
}

// Fulfills returns true the provided user has an AuthLevel >= the conditions level
func (c InSets) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(*v0.User)
	for _, n := range c.sets {
		if n == s.Set {
			return ok
		}
	}
	return false
}

// GetName returns the UserAuthGreater condition's name.
func (c InSets) GetName() string {
	return "InSets"
}

func newInSets(any *any.Any) (ladon.Condition, error) {
	arg := &v0.InSetsArg{}
	err := ptypes.UnmarshalAny(any, arg)
	if err != nil {
		return nil, err
	}
	return InSets{sets: arg.Sets}, nil
}
