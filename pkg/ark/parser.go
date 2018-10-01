package ark

import (
	"encoding/json"
	"io"

	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
)

func (s Service) parse(reader io.Reader) (tx *ArkV2.Transaction, err error) {
	tx = &ArkV2.Transaction{}
	return tx, json.NewDecoder(reader).Decode(tx)
}
