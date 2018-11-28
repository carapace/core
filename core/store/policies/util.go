package policies

import (
	"github.com/carapace/core/api/v0/proto"
	"strings"
)

var separators = []string{
	"//", "++", "##",
}

func SerializeActions(actions ...v0.Action) string {
	for _, sep := range separators {
		for _, action := range actions {
			if !strings.Contains(action.String(), sep) {
				return serialize(sep, actions...)
			}
		}
	}
	panic("unserializable action")
}

func DeserializeActions(actions string) []v0.Action {
	sep := actions[0:2]
	items := strings.Split(actions, sep)[1:]
	items = items[:len(items)-1]

	var res []v0.Action

	for _, item := range items {
		enum, ok := v0.Action_value[item]
		if !ok {
			panic("unable to deserialize actions")
		}
		res = append(res, v0.Action(enum))
	}
	return res
}

func serialize(separator string, actions ...v0.Action) string {
	res := separator
	for _, action := range actions {
		res = res + action.String() + separator
	}
	return res
}
