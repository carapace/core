package ark

import (
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/gin-gonic/gin/json"
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
