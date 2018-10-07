package test

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/dchest/uniuri"
)

func GenSecret(asset v1.Asset, namespace v1.Namespace) v1.Secret {
	s := uniuri.New()
	s = s + asset.String()
	s2 := s + namespace.String()

	return v1.Secret{
		Secrets:   append([][]byte{[]byte(s), []byte(s2)}),
		Asset:     asset,
		Namescace: &namespace,
	}
}
