package store

import (
	"database/sql"
	"github.com/carapace/core/core/store/sets"
	"github.com/carapace/core/core/store/users"
	"github.com/pressly/goose"
)

type Manager struct {
	db *sql.DB

	Sets  *sets.Manager
	Users *user.Manager
}

func (m *Manager) Begin() (*sql.Tx, error) {
	return m.db.Begin()
}

func New(db *sql.DB) *Manager {
	return &Manager{
		db:    db,
		Sets:  sets.New(),
		Users: user.New(),
	}
}

func (m Manager) AutoMigrate() error {
	err := goose.SetDialect("sqlite3") // TODO make this database independent
	if err != nil {
		return err
	}
	return goose.Up(m.db, ".")
}
