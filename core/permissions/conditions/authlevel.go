package condition

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
)

// AuthLevelGTE validates that the auth level is GTE
type AuthLevelGTE struct {
	Level int32
}

// Fulfills returns true the provided user has an AuthLevel >= the conditions level
func (c AuthLevelGTE) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(*v0.User)
	return ok && s.AuthLevel >= c.Level
}

// GetName returns the UserAuthGreater condition's name.
func (c AuthLevelGTE) GetName() string {
	return "AuthLevelGTE"
}

func newUserAuthGreater(any *any.Any) (ladon.Condition, error) {
	arg := &v0.AuthLevelGreaterArg{}
	err := ptypes.UnmarshalAny(any, arg)
	if err != nil {
		return nil, err
	}
	return AuthLevelGTE{Level: arg.Level}, nil
}
