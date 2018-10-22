package append

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
)

var (
	KeyExists    = errors.New("key already exists")
	KeyNotExists = errors.New("key does not exist")
)

// Put is a convenience wrapper around New and Amend, creating a new entry in the DB either way, and returning the
// current state of the object.
func (db *DB) Put(key string, obj any.Any, meta any.Any) (chunk *Chunk, err error) {
	save := db.cache.Lock()
	defer db.cache.Unlock()
	defer save.Rollback()

	chunk, err = db.new(key, obj, meta)

	// either the key already exists or something else went wrong
	if err != nil {
		switch err {
		case KeyExists:
			chunk, err = db.amend(key, obj, meta)
			if err != nil {
				return nil, err
			}
			save.Commit()
			return
		default:
			return nil, err
		}
	}
	save.Commit()
	return
}

// New creates a new object. If a key already exists, an error is thrown.
func (db *DB) New(key string, obj any.Any, meta any.Any) (chunk *Chunk, err error) {
	save := db.cache.Lock()
	defer db.cache.Unlock()
	defer save.Rollback()

	chunk, err = db.new(key, obj, meta)
	if err != nil {
		return
	}

	// commit the changes to the cache
	save.Commit()
	return
}

// Amend creates a new amend to an existing object. If the key does not exist, an error is thrown.
func (db *DB) Amend(key string, obj any.Any, meta any.Any) (chunk *Chunk, err error) {
	save := db.cache.Lock()
	defer db.cache.Unlock()
	defer save.Rollback()

	chunk, err = db.amend(key, obj, meta)
	if err != nil {
		return
	}

	save.Commit()
	// get the newest object state
	return db.Get(key)
}

// createState fills all non-user provided fields related to DB state. It is required to take a lock
// on the DB cache before calling this function.
func (db *DB) createState(object Object) (chunk *Chunk, err error) {
	// Reset ensures our result object is properly constructed
	chunk.Reset()

	// Instantiating the object fields
	chunk.Obj = &object
	chunk.State.ObjHash, err = db.hasher.Hash(object)
	if err != nil {
		return nil, err
	}

	// Getting the previous chain hash and combining it with the object hash
	chunk.State.ChainHash, err = db.hasher.CombineHash(chunk.State.ObjHash, db.cache.ChainHash())
	if err != nil {
		return nil, err
	}

	// Last field remaining is the Signature
	chunk.State.Signature, err = db.signer.Sign(chunk)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

// A wrapper around cellar.append and proto.Marshal
func (db *DB) append(chunk proto.Message) (pos int64, err error) {
	// marshalling to bytes
	var btes []byte
	btes, err = proto.Marshal(chunk)
	if err != nil {
		return
	}

	// appending to the DB
	pos, err = db.cellar.Append(btes)
	return
}

// new is the unexported business logic of New, without any locking.
func (db *DB) new(key string, obj any.Any, meta any.Any) (chunk *Chunk, err error) {
	if db.cache.KeyExists(key) {
		return nil, KeyExists
	}

	db.cache.AddKey(key)

	// raw object, which is unaware of it's state
	o := Object{Key: key, Value: &obj, Meta: &Meta{Fields: &meta, Timestamp: ptypes.TimestampNow(), Type: DataType_Create}}

	// filling in all the remaining fields
	chunk, err = db.createState(o)
	if err != nil {
		return nil, err
	}

	_, err = db.append(chunk)
	if err != nil {
		return
	}

	// updating the cache
	db.cache.SetChainHash(chunk.State.ChainHash)
	return
}

func (db *DB) amend(key string, obj any.Any, meta any.Any) (chunk *Chunk, err error) {
	// First we check if the key exists
	if !db.cache.KeyExists(key) {
		return nil, KeyNotExists
	}

	// create the new amend
	// raw object, which is unaware of it's state
	o := Object{Key: key, Value: &obj, Meta: &Meta{Fields: &meta, Timestamp: ptypes.TimestampNow(), Type: DataType_Amend}}
	chunk, err = db.createState(o)
	if err != nil {
		return nil, err
	}

	_, err = db.append(chunk)
	if err != nil {
		return nil, err
	}

	db.cache.SetChainHash(chunk.State.ChainHash)
	return
}
