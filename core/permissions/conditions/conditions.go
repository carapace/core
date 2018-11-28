package condition

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ory/ladon"
)

var Factory = map[string]func(any *any.Any) (ladon.Condition, error){
	v0.ConditionNames_AuthLevelGTE.String(): newUserAuthGreater,
	v0.ConditionNames_UsersOwns.String():    newUserOwns,
	v0.ConditionNames_InSets.String():       newInSets,
	v0.ConditionNames_MultiSig.String():     newMultisig,
}
