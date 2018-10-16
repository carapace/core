package test

import (
	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
	"testing"
	"time"
)

func TestBolt(t *testing.T) {
	t.Parallel()

	db, exit := Bolt(0600, &bolt.Options{Timeout: 1 * time.Second})
	defer exit()
	assert.NoError(t, db.Sync())
}
