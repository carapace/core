package install

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/pkg/router"
	"github.com/carapace/core/pkg/v0"
	"github.com/carapace/core/pkg/v0/ownerset"
)

func init() {
	// install the handlers with the v0 router
	router.Router.RegisterInfoService(v0_handler.Info{
		Info: &v0.Info{
			ApiVersion:  "v0",
			Semantic:    "v0.0.0",
			Kinds:       []string{"none"},
			Mode:        v0.Mode_Debug,
			Description: "development bootstrap of carapace",
		},
	})

	router.Router.Register("v0", "ownerSet", &ownerset.Handler{})

}
