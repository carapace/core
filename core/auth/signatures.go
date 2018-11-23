//go:generate mockgen -destination=mocks/signer_mock.go -package=mock github.com/carapace/core/core/auth Signer

package auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/proto"
)

type Signer interface {
	Check(pubkey *ecdsa.PublicKey, message proto.Message, signature *v0.Signature) (bool, error)
	Sign(key *ecdsa.PrivateKey, obj proto.Message) (*v0.Signature, error)
}

type DefaultSigner struct{}

func (d DefaultSigner) Sign(key *ecdsa.PrivateKey, obj proto.Message) (*v0.Signature, error) {
	bytes, err := proto.Marshal(obj)
	if err != nil {
		return nil, err
	}
	r, s, err := ecdsa.Sign(rand.Reader, key, bytes)
	if err != nil {
		return nil, err
	}

	return &v0.Signature{
		R: r.Bytes(),
		S: s.Bytes(),
	}, nil
}

func (d DefaultSigner) Check(pubkey *ecdsa.PublicKey, message proto.Message, signature *v0.Signature) (bool, error) {
	var r = &big.Int{}
	var s = &big.Int{}

	r.SetBytes(signature.R)
	s.SetBytes(signature.S)

	bytes, err := proto.Marshal(message)
	if err != nil {
		return false, err
	}

	return ecdsa.Verify(pubkey, bytes, r, s), nil
}
