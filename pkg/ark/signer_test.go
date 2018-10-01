package ark

import (
	"testing"

	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Sign2(t *testing.T) {
	t.Parallel()
	tx := &ArkV2.Transaction{
		Id:          "dummy",
		Type:        0,
		Amount:      10000000,
		Fee:         10000000,
		Signature:   "dummy",
		VendorField: "dummy",
		Timestamp:   ArkV2.GetTime(),
	}

	newService().sign(tx)

	correct, err := tx.Verify() // Verify is a function provided by ArkEcosystem/go-crypto to validate a transaction
	require.NoError(t, err)
	assert.True(t, correct)
}
