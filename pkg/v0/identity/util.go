package identity

import (
	"github.com/carapace/core/api/v0/proto"
)

func GetMaxAuthFromAccess(access []*v0.AccessProtocol) int32 {
	var max int32
	for _, protocol := range access {
		if protocol.GetAuthLevel() > max {
			max = protocol.GetAuthLevel()
		}
	}
	return max
}

func GetUserAccessSet(access []*v0.AccessProtocol) map[string]struct{} {
	set := make(map[string]struct{})
	for _, protocol := range access {
		set[protocol.GetUser()] = struct{}{}
	}
	return set
}

func GetUserSetAccessSet(access []*v0.AccessProtocol) map[string]struct{} {
	set := make(map[string]struct{})
	for _, protocol := range access {
		set[protocol.GetUserSet()] = struct{}{}
	}
	return set
}
