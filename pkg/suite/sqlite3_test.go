package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlite3(t *testing.T) {
	assert.NotPanics(t, func() {
		db, exit := Sqlite3(t)
		defer exit()

		assert.NotNil(t, db)
		assert.NoError(t, db.Ping())
	})
}
