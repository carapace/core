package condition

import (
	"bytes"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
)

// UserOwns is an exemplary condition.
type UserOwns struct {
	names []string
	keys  [][]byte
}

// Fulfills returns true the provided user has an AuthLevel >= the conditions level
func (c UserOwns) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(*v0.User)

	for _, n := range c.names {
		if n == s.Name {
			return ok
		}
	}

	for _, n := range c.keys {
		if bytes.Equal(n, s.PrimaryPublicKey) {
			return ok
		}
	}

	return false
}

// GetName returns the UserAuthGreater condition's name.
func (c UserOwns) GetName() string {
	return "UserOwns"
}

func newUserOwns(any *any.Any) (ladon.Condition, error) {
	arg := &v0.UserOwnsArg{}
	err := ptypes.UnmarshalAny(any, arg)
	if err != nil {
		return nil, err
	}
	return UserOwns{names: arg.Names, keys: arg.PrimaryPublicKeys}, nil
}
