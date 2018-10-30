package test

import (
	"github.com/carapace/core/pkg/chain-db"
	"github.com/stretchr/testify/require"
	"testing"
)

func ChainDB(t *testing.T, config chaindb.Config, option ...chaindb.ConfOption) (*chaindb.DB, func()) {
	var exit func()
	config.Folder, exit = Dir(t, string(hash(caller())))

	db, err := chaindb.New(config, option...)
	require.NoError(t, err)
	return db, func() {
		db.Close()
		exit()
	}
}
