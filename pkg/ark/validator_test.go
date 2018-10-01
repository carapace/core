package ark

import (
	"testing"

	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/stretchr/testify/assert"
)

// TestValidator_Validate_TXValidators checks the default middlewares which verify
// if a transaction contains all necessary fields
func TestValidator_Validate_TXValidators(t *testing.T) {
	tcs := []struct {
		tx *ArkV2.Transaction

		wantErr bool
		desc    string
	}{
		{
			tx: &ArkV2.Transaction{
				Id:          "dummy",
				Type:        0,
				Amount:      10000000,
				Fee:         10000000,
				Signature:   "dummy",
				VendorField: "dummy",
				Timestamp:   ArkV2.GetTime(),
			},
			wantErr: false,
			desc:    "a transaction where all fields are set should pass",
		},
		{
			tx: &ArkV2.Transaction{
				Id:          "dummy",
				Type:        100,
				Amount:      10000000,
				Fee:         10000000,
				Signature:   "dummy",
				VendorField: "dummy",
				Timestamp:   ArkV2.GetTime(),
			},

			wantErr: true,
			desc:    "a transaction where ID is out of range should fail",
		},
		{
			tx: &ArkV2.Transaction{
				Id:          "dummy",
				Type:        0,
				Amount:      0,
				Fee:         10000000,
				Signature:   "dummy",
				VendorField: "dummy",
				Timestamp:   ArkV2.GetTime(),
			},

			wantErr: true,
			desc:    "a transaction where Amount is 0 should fail",
		},
		{
			tx: &ArkV2.Transaction{
				Id:          "dummy",
				Type:        0,
				Amount:      10000000,
				Fee:         0,
				Signature:   "dummy",
				VendorField: "dummy",
				Timestamp:   ArkV2.GetTime(),
			},

			wantErr: true,
			desc:    "a transaction where fee is 0 should fail",
		},
	}

	val := Validator{
		txmiddleware: []TXValidationfunc{
			typeIsSet,
			amountIsSet,
			feeIsSet,
		},
	}

	for _, tc := range tcs {
		err := val.validateTX(tc.tx)
		if tc.wantErr {
			assert.Error(t, err, tc.desc)
		} else {
			assert.NoError(t, err, tc.desc)
		}
	}
}
