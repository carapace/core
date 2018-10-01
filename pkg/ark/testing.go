package ark

import (
	"encoding/json"
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
)

func newService() *Service {
	return &Service{}
}

func newTx() *ArkV2.Transaction {
	return &ArkV2.Transaction{
		Id:          "dummy",
		Type:        0,
		Amount:      10000000,
		Fee:         10000000,
		VendorField: "dummy",
		Timestamp:   ArkV2.GetTime(),
	}
}

func newJSONtx() ([]byte, error) {
	return json.Marshal(newTx())
}
