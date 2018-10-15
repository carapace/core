package core

import (
	"context"

	"github.com/carapace/core/api/v1/proto/generated"
)

// Config is the endpoint defined in api/proto/v1/services.proto. It is used to pass
// human defined configuration files.
func (c *Core) Config(ctx context.Context, conf *v1.Config) (*v1.Response, error) {
	r, err := c.Conf.In.In(ctx, *conf)
	if err != nil {
		return ResponseFromErr(err), nil
	}
	return ResponseSuccess(r), nil
}
