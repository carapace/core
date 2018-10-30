package test

import (
	"github.com/carapace/cellar"
	"github.com/carapace/core/pkg/chain-db"
	"github.com/carapace/core/pkg/state"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChainDB(t *testing.T) {

	folder, df := Dir(t)
	defer df()

	// don't copy this config for production usage
	c := chaindb.Config{
		Folder:        folder,
		Hasher:        state.EasyHasher{},
		Signer:        state.NewHMAC([]byte("")),
		Verifier:      state.NewHMAC([]byte("")),
		Cache:         chaindb.NewMemCache(),
		CellarOptions: []cellar.Option{cellar.WithNoFileLock},
	}

	require.NoError(t, c.Build())

	_, exit := ChainDB(t, c)
	defer exit()
}
