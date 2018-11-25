package permissions

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/ory/ladon"
)

var Effects = map[v0.Effect]string{
	v0.Effect_Allow: ladon.AllowAccess,
	v0.Effect_Deny:  ladon.DenyAccess,
}

var Actions = map[v0.Action]string{
	v0.Action_Alter:         "alter",
	v0.Action_Use:           "use",
	v0.Action_EnableDisable: "EnableDisable",
}
