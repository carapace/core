package append

import (
	"github.com/carapace/cellar"
	"github.com/golang/protobuf/proto"
	"github.com/imdario/mergo"
)

type ReadOp func(chunk Chunk) error

// Read is the most low level read API exposed by Append DB. In general it should not be used.
func (db *DB) Read(op ReadOp) error {
	reader := db.cellar.Reader()

	err := reader.Scan(func(pos *cellar.ReaderInfo, data []byte) error {
		chunk := &Chunk{}

		err := proto.Unmarshal(data, chunk)
		if err != nil {
			return err
		}

		err = op(*chunk)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetChunks returns all chunks matching the key.
func (db *DB) GetChunks(key string) (res []Chunk, err error) {
	err = db.Read(func(chunk Chunk) error {
		if chunk.Obj.Key == key {
			res = append(res, chunk)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

// Get is an easy wrapper around mergo and db.GetChunks, returing the full state of an object
func (db *DB) Get(key string) (*Chunk, error) {
	chunks, err := db.GetChunks(key)
	if err != nil {
		return nil, err
	}

	res := &Chunk{}

	for _, c := range chunks {
		err = mergo.Merge(res, c, mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
