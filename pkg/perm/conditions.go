package perm

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
)

var ConditionsFactory = map[string]func(any *any.Any) (ladon.Condition, error){
	v0.ConditionNames_AuthLevelGreater.String(): newUserAuthGreater,
}

// UserAuthGreater is an exemplary condition.
type AuthLevelGreater struct {
	Level int32
}

// Fulfills returns true the provided user has an AuthLevel >= the conditions level
func (c AuthLevelGreater) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(*v0.User)
	return ok && s.AuthLevel >= c.Level
}

// GetName returns the UserAuthGreater condition's name.
func (c AuthLevelGreater) GetName() string {
	return "AuthLevelGreater"
}

func newUserAuthGreater(any *any.Any) (ladon.Condition, error) {
	arg := &v0.AuthLevelGreaterArg{}
	err := ptypes.UnmarshalAny(any, arg)
	if err != nil {
		return nil, err
	}
	return AuthLevelGreater{Level: arg.Level}, nil
}
