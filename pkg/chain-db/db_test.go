package chaindb

import (
	"fmt"
	"hash/fnv"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/carapace/core/pkg/chain-db/sqlite3"
	"github.com/carapace/core/pkg/suite"

	"github.com/boltdb/bolt"
	"github.com/carapace/cellar"
	"github.com/carapace/core/pkg/state"
	"github.com/stretchr/testify/require"
)

const (
	defaultSecret = "defaultSecret"
)

// small helper function to hash strings to generate unique names
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

var gonce = &sync.Once{}
var c Config

func getConf(t *testing.T) Config {
	gonce.Do(func() {
		newConf(t)
	})
	return c
}

func newConf(t *testing.T) Config {
	blt, err := bolt.Open(
		fmt.Sprintf("./%s.%s", string(hash(caller())), "bolt"),
		0600,
		&bolt.Options{Timeout: 2 * time.Second},
	)
	require.NoError(t, err)
	metadb := &cellar.BoltMetaDB{DB: blt}
	metadb.Init()

	c = Config{
		Folder:   ".",
		Hasher:   state.EasyHasher{},
		Signer:   state.NewHMAC([]byte(defaultSecret)),
		Verifier: state.NewHMAC([]byte(defaultSecret)),
		Cache:    NewMemCache(),
		MetaDB:   metadb,
		Logger:   defaultLogger(),
	}
	require.NoError(t, c.Build())
	return c
}

func getDB(t *testing.T) (*DB, func()) {
	sql, exit := test.Sqlite3(t)
	db := newDB(getConf(t), t)
	store := sqlite3.New(sql)
	store.Migrate()
	db.config.Store = store
	return db, func() {
		exit()
	}
}

func newDB(conf Config, t *testing.T) *DB {
	var err error
	db, err := New(conf)
	require.NoError(t, err)
	return db
}

// caller returns the full function which called this function.
func caller() string {
	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "n/a"
	}

	// return its name
	return fun.Name()
}

func TestConfig_Build(t *testing.T) {
	getConf(t)
}

func TestNew(t *testing.T) {
	newDB(getConf(t), t)
}
