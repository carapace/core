package ark

import (
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/gin-gonic/gin/json"
)

const (
	transferTX = 0
)

func (s Service) createTX(params *v1.Transaction) (*v1.TransactionResponse, error) {
	tx := &ArkV2.Transaction{
		Type:      transferTX,
		Amount:    uint64(params.Amount),
		Fee:       uint64(params.Fee),
		Timestamp: ArkV2.GetTime(),
	}

	s.sign(tx)
	payload, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}

	return &v1.TransactionResponse{
		Payload: payload,
	}, nil
}
