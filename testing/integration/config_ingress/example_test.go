package config_ingress_test

import (
	"context"
	"fmt"

	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/core/options"
	"github.com/golang/protobuf/ptypes"

	"github.com/carapace/core/internal/core"
	"github.com/carapace/core/pkg/handlers/mock"

	// installing the mock config ingress handler. This registers the message type
	// Mock in the scheme, the validator to check the message, and the handler to
	// return a response.
	_ "github.com/carapace/core/pkg/handlers/mock/install"
)

func Example() {
	// Initialize the inner app. Normally we'd pass this app to a grpcServer
	c := core.New(&core.Config{},
		options.WithMock,
	)

	// the client creates a message type
	spec, _ := ptypes.MarshalAny(
		&mock.Mock{
			Mock: "true",
		})
	message := v1.Config{
		Header: &v1.Header{
			ApiVersion: mock.Version,
			Kind:       mock.Kind,
		},
		Spec: spec,
	}

	// normally we'd use a grpc client to pass the message
	res, err := c.Config(context.Background(), &message)
	if err != nil {
		panic("example went wrong!  " + err.Error())
	}

	// we are returned the following:
	fmt.Println(res.MSG)
	fmt.Println(res.Err)
	fmt.Println(res.Code)
	// Output:
	// mocking!
	//
	// OK
}
