package bitcoin

import (
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/pkg/errors"
)

type Validator struct {
	txmiddleware []TXValidationfunc
}

func (v *Validator) Validate(tx *ArkV2.Transaction) error {
	return v.validateTX(tx)
}

func (v *Validator) validateTX(tx *ArkV2.Transaction) error {
	for _, f := range v.txmiddleware {
		err := f(tx)
		if err != nil {
			return err
		}
	}
	return nil
}

type TXValidationfunc func(tx *ArkV2.Transaction) error

func typeIsSet(tx *ArkV2.Transaction) error {
	if tx.Type > 4 {
		return errors.New("transaction type should be set")
	}
	return nil
}

func amountIsSet(tx *ArkV2.Transaction) error {
	if tx.Amount <= 0 {
		return errors.New("amount should be set")
	}
	return nil
}

func feeIsSet(tx *ArkV2.Transaction) error {
	if tx.Fee <= 0 {
		return errors.New("fee should be set")
	}
	return nil
}

func newValidator() *Validator {
	return &Validator{
		txmiddleware: []TXValidationfunc{},
	}
}
