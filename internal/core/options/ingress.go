package options

import (
	"github.com/carapace/core/internal/core"
	"github.com/carapace/core/internal/ingress"
	"github.com/carapace/core/internal/store/mock"
)

// compile time assertion to check interface
var _ core.Option = WithMock

func WithMock(c *core.Core) {
	c.Conf.In = &ingress.DefaultConfig{Store: &mock.Store{}}
}
