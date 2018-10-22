package state

import (
	"github.com/mitchellh/hashstructure"
)

// Hasher is a common interface for a generic hasher.
type Hasher interface {
	Hash(interface{}) (uint64, error)
	CombineHash(...interface{}) (uint64, error)
}

// EasyHasher implements a simple Hasher using hashstructure
type EasyHasher struct{}

// Hash returns the hash of an object
func (e EasyHasher) Hash(i interface{}) (uint64, error) {
	return hashstructure.Hash(i, nil)
}

// CombineHash first hashes all of the objects in order, storing the results in a slice, then hashing each hash with
// the previous combination of two hashes.
//
// the first combination hash results from hashing hashes[0] with the nil value of uint64 (0).
func (e EasyHasher) CombineHash(objs ...interface{}) (uint64, error) {
	var hashes []uint64

	for _, obj := range objs {
		hash, err := e.Hash(obj)
		if err != nil {
			return 0, err
		}
		hashes = append(hashes, hash)
	}

	var hash uint64
	for _, h := range hashes {
		var err error
		hash, err = e.Hash(hash + h)
		if err != nil {
			return 0, err
		}
	}
	return hash, nil
}
