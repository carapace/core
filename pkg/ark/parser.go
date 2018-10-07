package ark

import (
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/carapace/core/api/v1/proto/generated"
)

func Parse(transaction v1.Transaction) ArkV2.Transaction {
	var txType byte
	switch transaction.Type {
	case v1.TransactionType_Transfer:
		txType = ArkV2.TRANSACTION_TYPES.Transfer
	case v1.TransactionType_Vote:
		txType = ArkV2.TRANSACTION_TYPES.Vote
	default:
		panic("unrecognised tx type")
	}

	return ArkV2.Transaction{
		Amount:      uint64(transaction.Amount),
		Fee:         uint64(transaction.Fee),
		Timestamp:   ArkV2.GetTime(),
		RecipientId: transaction.Recipient,
		VendorField: string(transaction.Meta),
		Type:        txType,
	}
}
