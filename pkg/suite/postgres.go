package test

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// Postgres returns a db connection to a clean postgres DB, and a cleanup function which should be deferred
func Postgres() (db *sql.DB, cleanup func()) {
	Logger.Debug("creating postgres instance")

	decrement := resourceLimit.allow()

	resource, err := pool.Run("circleci/postgres", "10", []string{"POSTGRES_USER=root", "POSTGRES_DB=db", "POSTGRES_PASSWORD=secret"})
	if err != nil {
		Logger.Panic("Could not start postgres resource:", zap.Error(err))
	}

	Logger.Debug("created postgres resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		Logger.Debug("connecting to postgres resource")

		db, err = sql.Open("postgres", fmt.Sprintf("postgres://root:secret@localhost:%s/db?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			Logger.Warn("error while dialing postgres", zap.Error(err))
			return err
		}
		return db.Ping()
	})
	Logger.Debug("connected to postgres resource")

	if err != nil {
		Logger.Panic("Could not connect to postgres docker image", zap.Error(err))
	}
	return db, func() {
		err := db.Close()
		if err != nil {
			Logger.Warn("unable to close postgres", zap.Error(err))
		}
		err = resource.Close()
		if err != nil {
			Logger.Warn("unable to close postgres resource", zap.Error(err))
		}

		decrement()
	}
}
