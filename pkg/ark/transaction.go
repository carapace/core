package ark

import (
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/pkg/errors"
)

const (
	transferTX = 0
)

func (s Service) createTX(params *v1.Transaction, secrets v1.Secret) (*ArkV2.Transaction, error) {
	tx := Parse(*params)

	switch len(secrets.Secrets) {
	case 0:
		return nil, errors.New("wallet does not have generated Ark addresses")
	case 1:
		tx.Sign(string(secrets.Secrets[0]))
		return &tx, nil
	case 2:
		tx.Sign(string(secrets.Secrets[0]))
		tx.SecondSign(string(secrets.Secrets[1]))
		return &tx, nil
	default:
		return nil, errors.New("Incorrect amount of secrets provided")
	}
}
