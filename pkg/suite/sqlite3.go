package test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func Sqlite3(t *testing.T, dir ...string) (*sql.DB, func()) {
	loc := path.Join(".", string(hash(caller())))

	if len(dir) == 1 {
		loc = dir[0]
	}

	db, err := sql.Open("sqlite3", loc)
	require.NoError(t, err)

	return db, func() {
		db.Close()
		os.Remove(loc)
		os.Remove(fmt.Sprintf("%s-journal", loc))
	}
}
