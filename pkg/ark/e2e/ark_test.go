package e2e

import (
	"bytes"
	"testing"

	"encoding/json"
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/carapace/core/pkg/ark"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_E2E(t *testing.T) {
	s := ark.New(ark.Config{}, ark.WithDefault)

	tx := &ArkV2.Transaction{
		Id:          "dummy",
		Type:        0,
		Amount:      10000000,
		Fee:         10000000,
		VendorField: "dummy",
		Timestamp:   ArkV2.GetTime(),
	}

	payload, err := json.Marshal(tx)
	require.NoError(t, err)

	err = s.Validate(bytes.NewReader(payload))
	require.NoError(t, err)

	signed, err := s.Sign(bytes.NewReader(payload))
	require.NoError(t, err)

	ok, err := s.Verify(signed)
	require.NoError(t, err)
	assert.True(t, ok)
}
