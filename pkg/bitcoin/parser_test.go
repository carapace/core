package bitcoin

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_parse(t *testing.T) {
	t.Parallel()

	want := &Transaction{
		Amount:             10000000,
		TxId:               "hbjnkjiuhgyhbnjiughbn",
		SourceAddress:      "dummy",
		DestinationAddress: "dummy",
		UnsignedTx:         "dummy",
	}

	js, err := json.Marshal(want)
	require.NoError(t, err)

	got, err := newService().parse(bytes.NewReader(js))
	require.NoError(t, err)

	assert.EqualValues(t, got, want)
}
