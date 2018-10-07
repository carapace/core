package test

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

func GetNamespace(seed int) (space v1.Namespace) {
	if seed == 1 {
		space = v1.Namespace{
			Region: "Netherlands",
			Name:   "MyWallet",
		}
	} else {
		space = v1.Namespace{
			Region: "Netherlands" + string(seed),
			Name:   "MyWallet" + string(seed),
		}
	}
	return space
}
