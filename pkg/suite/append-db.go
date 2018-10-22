package test

import (
	"github.com/carapace/core/pkg/append-db"
	"os"
)

func AppendDB(options ...append.Option) (*append.DB, func()) {
	// Create folder if it does not exist
	if _, err := os.Stat(string(hash(caller()))); os.IsNotExist(err) {
		os.Mkdir(string(hash(caller())), os.ModePerm)
	}

	db, err := append.New(string(hash(caller())), options...)
	if err != nil {
		panic("testsuite: db")
	}
	return db, func() {
		db.Close()
		os.Remove(string(hash(caller())))
	}
}
