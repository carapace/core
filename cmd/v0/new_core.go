package main

import (
	"github.com/carapace/core/pkg/v0/userset"
	"os"
	"testing"

	"github.com/carapace/core/core"
	"github.com/carapace/core/core/auth"
	"github.com/carapace/core/core/router"
	"github.com/carapace/core/pkg/health"
	"github.com/carapace/core/pkg/suite"
	"github.com/carapace/core/pkg/v0/ownerset"
	"github.com/stretchr/testify/require"
)

func (s *Suite) NewCore(t *testing.T, rootDir string) (*core.App, func()) {

	db, dbCleanup := test.Sqlite3(t, rootDir)
	store, err := core.NewStore(db)
	require.NoError(t, err)

	cfg, err := core.Config{
		Router:        router.New(),
		Store:         store,
		HealthManager: health.New(),
		Health: core.Health{
			Port: "9000",
			Host: "0.0.0.0",
		},
	}.Build()
	require.NoError(t, err)

	app, err := core.New(cfg)
	require.NoError(t, err)

	// install the handlers with the v0 router
	authManager := &auth.Manager{
		Signer:  auth.DefaultSigner{},
		Store:   app.Store,
		Decoder: auth.X509Marshaller{},
	}

	app.Router.Register(
		authManager.RootOrBackupOrNoOwners(ownerset.New(authManager, store)),
		authManager.RegularAuth(1, 1, 1, userset.New(store)),
	)

	return app, func() {
		dbCleanup()
		os.Remove(rootDir)
	}
}
