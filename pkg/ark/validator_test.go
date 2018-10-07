package ark

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestValidator_Validate_TXValidators checks the default middlewares which verify
// if a transaction contains all necessary fields
func TestValidator_Validate_TXValidators(t *testing.T) {
	tcs := []struct {
		tx *v1.Transaction

		wantErr bool
		desc    string
	}{
		{
			tx: &v1.Transaction{
				Amount:    10000000,
				Fee:       10000000,
				Recipient: "APcgT8sfTWkBXb4xLMjyEhJrga1CcCFV3Z",
				Type:      v1.TransactionType_Transfer,
			},
			wantErr: false,
			desc:    "a transaction where all fields are set should pass",
		},
		{
			tx: &v1.Transaction{
				Amount: 0,
				Fee:    10000000,
				Type:   v1.TransactionType_Transfer,
			},

			wantErr: true,
			desc:    "a transaction where Amount is 0 should fail",
		},
		{
			tx: &v1.Transaction{
				Amount: 10000000,
				Fee:    0,
				Type:   v1.TransactionType_Transfer,
			},

			wantErr: true,
			desc:    "a transaction where fee is 0 should fail",
		},
		{
			tx: &v1.Transaction{
				Amount: 10000000,
				Fee:    1,
				Type:   v1.TransactionType_Vote,
			},

			wantErr: false,
			desc:    "a transaction where type is vote should pass",
		},
		{
			tx: &v1.Transaction{
				Amount: 10000000,
				Fee:    1,
				Type:   v1.TransactionType_SCCreation,
			},

			wantErr: true,
			desc:    "a transaction where type is TransactionType_SCCreation should fail",
		},
		{
			tx: &v1.Transaction{
				Amount: 10000000,
				Fee:    1,
				Type:   v1.TransactionType_SCInvocation,
			},

			wantErr: true,
			desc:    "a transaction where type is TransactionType_SCInvocation should fail",
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
		err := val.validateTX(*tc.tx)
		if tc.wantErr {
			assert.Error(t, err, tc.desc)
		} else {
			assert.NoError(t, err, tc.desc)
		}
	}
}
