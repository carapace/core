# Config Handlers

A config handler is a struct fulfilling the interface: core.ConfigHandler
defined in internal/core/handler.go A handler is called after the new configuration
has been committed in the document store.

## How to write a handler
When writing a handler, first the ApiVersion and Kind needs to be determined.
Then the handler interface needs to be fulfilled. Using protobuf, a message spec
needs to be defined. Finally a validation function for the message needs to be
defined.

### Example
```go
/*
package mock defines a mock configHandler
*/
package mock


# The version and kind used to identify the correct handler
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
func (h *Handler) Call(config v1.Config) (result core.Response, err error) {
	return &Response{msg: "mocking!", err: ""}, nil
}

// Validate checks that only the field mock is present in the spec
func Validate(config v1.Config) error {
	configObj := &Mock{} # struct Mock is defined in mock.proto
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

### mock.proto

syntax="proto3";
package mock;
option go_package = "mock";

message Mock {
    string Mock = 1;
}

### install.go

// register the handler in our API scheme.
func init() {
	scheme.Register(Version, Kind, &Handler{}, Validate)
}
```
