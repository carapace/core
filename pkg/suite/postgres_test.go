package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	db, exit := Postgres()
	defer exit()
	assert.NoError(t, db.Ping())
}
