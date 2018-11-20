package test

import (
	"fmt"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"hash/fnv"
	"os"
)

// small helper function to hash strings to generate unique names
func hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	return h.Sum32()
}

// Returns a connection to a clean boltdb in RW mode, and a function to clean up the resource
func Bolt(mode os.FileMode, opt *bolt.Options) (db *bolt.DB, cleanup func()) {
	var err error
	db, err = bolt.Open(fmt.Sprintf("%s.db", string(hash(caller()))), mode, opt)
	if err != nil {
		Logger.Panic("error opening boltdb", zap.Error(err))
	}

	return db, func() {
		err = os.Remove(db.Path())
		if err != nil {
			Logger.Warn("unable to delete boltdb file", zap.Error(err), zap.String("PATH", db.Path()))
		}

		err := db.Close()
		if err != nil {
			Logger.Warn("unable to close boltdb", zap.Error(err))
		}
	}
}
