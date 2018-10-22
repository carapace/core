package proto_any_test

import (
	"fmt"

	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func exampleTrigger(config v1.Config) {
	configObj := &v1.WalletGen{}
	err := ptypes.UnmarshalAny(config.Spec, configObj)
	if err != nil {
		panic(err)
	}
	fmt.Println("I'm a walletGen object!", configObj.String())
}

func exampleClient() {
	configObj := &v1.WalletGen{Name: "myWallet"}
	path := proto.MessageName(configObj)
	fmt.Println("the helper func found this as my path: ", path)
	any, _ := ptypes.MarshalAny(configObj)
	fmt.Println("here is my actual defintion: ", any.TypeUrl)

	// post the message
	exampleTrigger(v1.Config{Spec: any})
}

func Example() {
	exampleClient()
	// Output:
	// the helper func found this as my path:  v1.config.walletGen
	// here is my actual defintion:  type.googleapis.com/v1.config.walletGen
	// I'm a walletGen object! Name:"myWallet"
}
