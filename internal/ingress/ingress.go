package ingress

import (
	"context"

	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core"
	"github.com/carapace/core/internal/scheme"
	"github.com/pkg/errors"
)

// DefaultConfig is an ingress controller.
type DefaultConfig struct {
	Store core.ConfigStore
}

func (d *DefaultConfig) In(ctx context.Context, conf v1.Config) (result string, err error) {
	err = d.validate(conf)
	if err != nil {
		return "", err
	}

	// point of no return; the config is being committed to the store
	res, err := d.Store.Add(conf)
	if err != nil {
		return "", err
	}

	// now we wait until the trigger has finished
	// TODO this call to Err blocks, use a select case to switch on ctx cancellation
	// the result has been committed up to now, so if the client quits on us, no need to
	// wait on the result
	if res.Err() != nil {
		return "", err
	}

	// return the result of the trigger
	return res.MSG(), nil
}

func (d *DefaultConfig) validate(conf v1.Config) error {
	// Check if the object is registered and validate it's contents
	err := scheme.Validate(conf)
	if err != nil {
		return errors.Wrap(err, "ingress:")
	}
	return nil
}
