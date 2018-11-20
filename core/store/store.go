package store

import (
	"database/sql"
	"github.com/carapace/core/core/store/sets"
	"github.com/carapace/core/core/store/users"
)

type Manager struct {
	db *sql.DB

	Sets  sets.Manager
	Users user.Manager
}

func (m *Manager) Begin() (*sql.Tx, error) {
	return m.db.Begin()
}

func New(db *sql.DB) (*Manager, error) {
	m := &Manager{db: db}

	tx, err := m.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = m.Sets.AutoMigrate(tx)
	if err != nil {
		return nil, err
	}

	err = m.Users.AutoMigrate(tx)
	if err != nil {
		return nil, err
	}
	return m, tx.Commit()
}
