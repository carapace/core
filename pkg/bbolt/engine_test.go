package bbolt

import (
	"github.com/carapace/core/pkg/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
	"testing"
	"time"
)

func TestEngine_Put(t *testing.T) {
	test.Strategy(t, test.LocalIntegration)

	b, exit := test.Bolt(0600, &bolt.Options{Timeout: 1 * time.Second})
	defer exit()
	engine := Engine{db: b}

	tcs := []struct {
		keys []string
		item []byte

		wantErr bool
		err     string
		desc    string
	}{
		{
			keys:    []string{"1", "2", "3"},
			item:    []byte("Hello"),
			wantErr: false,
		},
		{
			keys:    []string{"1", "2", "3"},
			item:    []byte("Hello"),
			wantErr: false,
		},
		{
			keys:    []string{"1", "", "3"},
			item:    []byte("Hello"),
			wantErr: false,
		},
		{
			keys:    []string{"", "", ""},
			item:    []byte("Hello"),
			wantErr: true,
			err:     ErrNoKeys.Error(),
			desc:    "not setting keys should return an error",
		},
		{
			keys:    []string{"1"},
			item:    []byte("Hello"),
			wantErr: false,
		},
		{
			keys:    []string{"1", "1", "2", "3", "1", "2", "3"},
			item:    []byte("Hello"),
			wantErr: false,
		},
	}

	for _, tc := range tcs {
		err := engine.Put(tc.item, tc.keys...)
		if tc.wantErr {
			assert.EqualError(t, err, tc.err, tc.desc)
			continue
		}
		assert.NoError(t, err, tc.desc)
	}
}

func TestEngine_GetAll(t *testing.T) {
	test.Strategy(t, test.LocalIntegration)

	b, exit := test.Bolt(0600, &bolt.Options{Timeout: 1 * time.Second})
	defer exit()

	engine := Engine{db: b}

	tcs := []struct {
		keys []string
		item []byte

		wantErr bool
		err     string
		desc    string
	}{
		{
			keys:    []string{"1", "2", "3"},
			item:    []byte("Hello"),
			wantErr: false,
		},
		{
			keys:    []string{"1", "1", "2", "3", "1", "2", "3"},
			item:    []byte("Hello"),
			wantErr: false,
		},
	}

	for _, tc := range tcs {
		err := engine.Put(tc.item, tc.keys...)
		if tc.wantErr {
			require.EqualError(t, err, tc.err, tc.desc)
			continue
		}
		require.NoError(t, err, tc.desc)
	}

	for _, tc := range tcs {
		item, err := engine.GetAll(tc.keys...)
		if tc.wantErr {
			continue
		}
		require.NoError(t, err, tc.desc)
		assert.EqualValues(t, tc.item, item[0])
	}
}
