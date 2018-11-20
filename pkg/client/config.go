package carapace

import (
	"crypto/ecdsa"
	"github.com/carapace/core/core/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Config struct {
	Credentials

	Host            string
	Port            string
	GRPC            *grpc.ClientConn
	PrivKey         *ecdsa.PrivateKey
	RecoveryPrivKey *ecdsa.PrivateKey
	Signer          auth.DefaultSigner
	Marshaller      auth.KeyMarshaller
}

type Credentials struct {
	Seed  string
	RSeed string
	Name  string
	Email string
}

func (c Config) Build() (*Config, error) {
	if c.Host == "" {
		return nil, errors.New("missing host")
	}
	if c.Port == "" {
		return nil, errors.New("missing port")
	}
	if c.GRPC == nil {
		return nil, errors.New("missing grpc client")
	}
	return &c, nil
}
