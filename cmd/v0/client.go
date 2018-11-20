package main

import (
	"github.com/carapace/core/core/auth"
	"github.com/carapace/core/pkg/client"
)

// NewClient returns a mock client without a grpc conn (it's API methods are all invalid)
func (s Suite) NewClient() *carapace.Client {
	cfg := carapace.Config{
		Marshaller: auth.X509Marshaller{},
		Signer:     auth.DefaultSigner{},
		Credentials: carapace.Credentials{
			Name:  "Karel",
			Email: "k.l.kubat@gmail.com",
			Seed:  "randomseed123",
			RSeed: "predeterminedseed321",
		},
	}
	return carapace.New(cfg)
}
