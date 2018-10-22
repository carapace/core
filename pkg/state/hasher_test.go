package state

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestEasyHasher_Hash(t *testing.T) {
	tcs := []struct {
		input  interface{}
		output uint64
	}{
		{
			input:  "teststringinput",
			output: 10911935989730201158,
		},
		{
			input:  "teststringinput2",
			output: 18222809183765790912,
		},
		{
			input:  "teststringinput3",
			output: 18222809183765790913, // found I hash collision I reckon :P // TODO check if hashtructure.Hash is secure
		},
		{
			input:  "teststringinput4",
			output: 18222809183765790918, // Same here... // TODO check if hashtructure.Hash is secure
		},
		{
			input:  18222809183,
			output: 16607209336448606860,
		},
		{
			input:  182228091831,
			output: 17539652398153896628,
		},
		{
			input:  1822280918311,
			output: 14496435365958069300,
		},
		{
			input: struct {
				Key string
				Val int
			}{
				Key: "key",
				Val: 1,
			},
			output: 428733854022535403,
		},
		{
			input: struct {
				Key string
				Val int
			}{
				Key: "erigsknjshilgnusl",
				Val: 18594257,
			},
			output: 12433465987827068531,
		},
	}

	h := EasyHasher{}

	for _, tc := range tcs {
		have, err := h.Hash(tc.input)
		require.NoError(t, err)
		assert.Equal(t, tc.output, have)
	}
}

func TestEasyHasher_CombineHash(t *testing.T) {
	tcs := []struct {
		input  []interface{}
		output uint64
	}{
		{
			input:  []interface{}{"val1", "val2", "val3"},
			output: 5512618098106155096,
		},
		{
			input:  []interface{}{"val1", 1, "val3"},
			output: 2399909538027278338,
		},
		{
			input:  []interface{}{"val1", 1, struct{ Key string }{"mykey"}},
			output: 6408917438915038148,
		},
	}

	h := EasyHasher{}

	for _, tc := range tcs {
		have, err := h.Hash(tc.input)
		require.NoError(t, err)
		assert.Equal(t, tc.output, have)
	}
}
