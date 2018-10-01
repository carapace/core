package ark

import (
	"bytes"
	"encoding/json"
	"testing"

	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_parse(t *testing.T) {
	t.Parallel()

	// Testdata is taken from the official Ark client at
	// https://github.com/ArkEcosystem/go-client/blob/master/client/two/transactions_test.go
	want := &ArkV2.Transaction{
		Id:          "dummy",
		Type:        0,
		Amount:      10000000,
		Fee:         10000000,
		Signature:   "dummy",
		VendorField: "dummy",
		Timestamp:   ArkV2.GetTime(),
	}

	js, err := json.Marshal(want)
	require.NoError(t, err)

	got, err := newService().parse(bytes.NewReader(js))
	require.NoError(t, err)

	assert.EqualValues(t, got, want)
}
