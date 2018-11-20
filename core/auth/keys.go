//go:generate mockgen -destination=mocks/keymarshaller_mock.go -package=mock github.com/carapace/core/pkg/v0/auth KeyMarshaller

package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
)

type KeyMarshaller interface {
	MarshalPublic(key *ecdsa.PublicKey) ([]byte, error)
	UnmarshalPublic(key []byte) (*ecdsa.PublicKey, error)

	MarshalPrivate(key *ecdsa.PrivateKey) ([]byte, error)
	UnmarshalPrivate(key []byte) (*ecdsa.PrivateKey, error)
}

// https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
type X509Marshaller struct{}

func (x X509Marshaller) MarshalPublic(key *ecdsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(key)
}

func (x X509Marshaller) UnmarshalPublic(key []byte) (*ecdsa.PublicKey, error) {
	genericPublicKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return nil, err
	}
	return genericPublicKey.(*ecdsa.PublicKey), nil
}

func (x X509Marshaller) MarshalPrivate(key *ecdsa.PrivateKey) ([]byte, error) {
	return x509.MarshalECPrivateKey(key)

}

func (x X509Marshaller) UnmarshalPrivate(key []byte) (*ecdsa.PrivateKey, error) {
	return x509.ParseECPrivateKey(key)
}
