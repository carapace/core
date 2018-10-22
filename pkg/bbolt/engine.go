package bbolt

import (
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

var (
	ErrNoKeys = errors.New("keys must be provided")
)

type Engine struct {
	db *bolt.DB
}

func (e *Engine) Put(item []byte, keys ...string) error {
	// remove unset keys
	keys = reduceKeys(func(s string) bool {
		return s != ""
	}, keys...)

	if len(keys) == 0 {
		return ErrNoKeys
	}

	tx, err := e.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback() // if we commit, this will return an error, no need to log it.

	var b *bolt.Bucket

	// opening the initial bucket
	bucket := tx.Bucket([]byte(keys[0]))
	if bucket != nil {
		b = bucket
	} else {
		b, err = tx.CreateBucketIfNotExists([]byte(keys[0]))
		if err != nil {
			return err
		}
	}

	// create a nested bucket structure for all keys
	for _, key := range keys[1:] {
		// check if the bucket already exists
		bucket = b.Bucket([]byte(key))
		if bucket != nil {
			b = bucket
			continue
		}
		b, err = b.CreateBucketIfNotExists([]byte(key))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	key, err := b.NextSequence()
	if err != nil {
		return err
	}

	// we store the object under an autoincrementing key. This allows for sorting later on.
	// TODO instead of incrementing uint64, use time.Time to audit when exactly the obj was stored.
	k := make([]byte, 8)
	binary.LittleEndian.PutUint64(k, uint64(key))
	err = b.Put(k, item)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) GetAll(keys ...string) ([][]byte, error) {
	// remove unset keys
	keys = reduceKeys(func(s string) bool {
		return s != ""
	}, keys...)

	if len(keys) == 0 {
		return nil, ErrNoKeys
	}
	var res [][]byte

	err := e.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(keys[0]))
		if b == nil {
			return errors.New("first key not found: " + keys[0])
		}

		// traverse the key to find lowest bucket
		for _, key := range keys[1:] {
			b = b.Bucket([]byte(key))
			if b == nil {
				return errors.New("first key not found: " + key)
			}
		}

		c := b.Cursor()
		// TODO I assume uint64 -> preserves order within bucket. This needs to be tested to make sure
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// check if val is a bucket
			if v == nil {
				continue
			}
			var cp = make([]byte, len(v))
			copy(cp, v)
			res = append(res, cp)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func reduceKeys(qualifier func(string) bool, keys ...string) []string {
	res := []string{}
	for _, k := range keys {
		if qualifier(k) {
			res = append(res, k)
		}
	}
	return res
}
