package state

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJSONEncoder_Encode_Decode(t *testing.T) {
	type TestStruct struct {
		Key string `json:"key"`
	}

	tcs := []struct {
		input  interface{}
		output []byte
		result interface{}
	}{
		{
			input:  "somedata1",
			output: nil,
			result: "",
		},
		{
			input:  float64(3456789),
			output: nil,
			result: float64(0),
		},
		{
			input:  float64(34522226789),
			output: nil,
			result: float64(0),
		},
		{
			input:  &TestStruct{Key: "mykey"},
			output: nil,
			result: &TestStruct{},
		},
	}

	encoder := JSONEncoder{}
	decoder := JSOnDecoder{}

	for _, tc := range tcs {
		have, err := encoder.Encode(tc.input)
		require.NoError(t, err)

		err = decoder.Decode(have, &tc.result)
		require.NoError(t, err)

		assert.Equal(t, tc.input, tc.result)
	}
}
