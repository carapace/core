package carapace

import (
	"github.com/carapace/core/core/auth"
	"path"
	"testing"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewClient() *Client {
	cfg := Config{
		Marshaller: auth.X509Marshaller{},
		Signer:     auth.DefaultSigner{},
		Credentials: Credentials{
			Name:  "Karel",
			Email: "k.l.kubat@gmail.com",
			Seed:  "somesecret1somesecret1",
			RSeed: "somesecret2somesecret2",
		},
	}
	return New(cfg)
}

func TestClient_LoadConfig(t *testing.T) {
	tcs := []struct {
		file             []string
		err              error
		desc             string
		expectedConfig   v0.Config
		expectedOwnerSet v0.OwnerSet
	}{
		{
			file: []string{"testdata", "ownerSet1.yaml"},
			err:  nil,
			desc: "simple ownerSet with single owner",

			expectedConfig: v0.Config{
				Header: &v0.Header{
					ApiVersion: "v0",
					Kind:       "type.googleapis.com/v0.OwnerSet",
				},
			},
			expectedOwnerSet: v0.OwnerSet{
				Quorum: 1,
				Owners: []*v0.User{
					{Name: "Karel", Email: "k.l.kubat@gmail.com", Weight: 1},
				},
			},
		},
		{
			file: []string{"testdata", "ownerSet3.yaml"},
			err:  nil,
			desc: "three owner-ownerSet with varying weight",

			expectedConfig: v0.Config{
				Header: &v0.Header{
					ApiVersion: "v0",
					Kind:       "type.googleapis.com/v0.OwnerSet",
				},
			},
			expectedOwnerSet: v0.OwnerSet{
				Quorum: 11,
				Owners: []*v0.User{
					{Name: "Karel", Email: "k.l.kubat@gmail.com", Weight: 10},
					{Name: "Robert", Email: "r.l.kubat@gmail.com", Weight: 1},
					{Name: "Laurens", Email: "l.l.kubat@gmail.com", Weight: 1},
				},
			},
		},
	}
	client := NewClient()

	for _, tc := range tcs {
		config, err := client.LoadConfig(path.Join(tc.file...))
		require.NoError(t, err)
		assert.Equal(t, tc.expectedConfig.Header.ApiVersion, config.Header.ApiVersion)
		assert.Equal(t, tc.expectedConfig.Header.Kind, config.Header.Kind)

		ownerset := &v0.OwnerSet{}
		require.NoError(t, ptypes.UnmarshalAny(config.Spec, ownerset))

		assert.Equal(t, tc.expectedOwnerSet.Quorum, ownerset.Quorum)

		for i, owner := range tc.expectedOwnerSet.Owners {
			assert.Equal(t, owner, ownerset.Owners[i])
		}
	}
}

func TestClient_LoadConfig_and_Sign(t *testing.T) {
	client := NewClient()
	require.NoError(t, client.GenPrivKey())
	config, err := client.LoadConfig(path.Join(path.Join("testdata", "ownerSet1.yaml")))
	require.NoError(t, err)

	err = client.SignConfig(config, nil)
	require.NoError(t, err)
}
