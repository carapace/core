package mock

import (
	"context"
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

const (
	Version = "v0"
	Kind    = "mock"
)

// Handler implements a very mock configuration handler, which only has one allowed
// field, which is mock, with value: "true"
//
// An example of a configuration file which is handled by the mockhandler is:
//
// ---
// apiVersion: v1
// kind: config
// metadata:
// 		name: whatever
// spec:
// 		mock: "true"
// ---
//
// calling this endpoint will cause the handler to return: "mocking!"
//
// implementing the mock handler causes a runtime warning message to be emitted
// during initialization.
//
// There are no security risks to this mock handler, it may
// be safely included in production code for e2e testing purposes.
type Handler struct {
}

// compile time check to verify handler matches the core.Handler interface
var _ core.ConfigHandler = &Handler{}

// Init does nothing, as the mock handler does not require any services
func (h *Handler) Init(core.Services) error {
	return nil
}

// Init does nothing, as the mock handler does not require any services
func (h *Handler) Call(ctx context.Context, config []*v1.Config, committer core.Committer) (result core.Response, err error) {
	defer committer.Commit()
	return &Response{msg: "mocking!", err: ""}, nil
}

// Validate checks that only the field mock is present in the spec
func Validate(config v1.Config) error {
	configObj := &Mock{}
	err := ptypes.UnmarshalAny(config.Spec, configObj)
	if err != nil {
		return err
	}
	if configObj.Mock != "true" {
		return errors.New(`mock handler: field Mock should be "true"`)
	}
	return nil
}

type Response struct {
	msg string
	err string
}

func (r *Response) MSG() string { return r.msg }
func (r *Response) Err() error  { return nil }
