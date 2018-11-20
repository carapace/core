package carapace

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"path"
	"testing"
)

func TestClient_LoadConfig_and_Add_User(t *testing.T) {
	client := NewClient()
	require.NoError(t, client.GenPrivKey())
	config, err := client.LoadConfig(path.Join(path.Join("testdata", "ownerSet1.yaml")))
	require.NoError(t, err)

	err = client.FillUserKeys(config)
	require.NoError(t, err)

	err = client.SignConfig(config, nil)
	require.NoError(t, err)

	fmt.Println(client.DumpString(config))
}
