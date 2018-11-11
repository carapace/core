package chaindb

import (
	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
)

// Option allows for further customization of Puts and Gets
type Option struct {
	Cached bool // allow the use of an internal memory cache, or completely reread from file
}

// defaultOpt returns the default behaviour when getting/putting
func defaultOpt() *Option {
	return &Option{Cached: false}
}

// Put enters a new item in the DB.
//
// Currently the user needs to provide if they are creating or amending an item; from the DB's perspective;
// this should not be necessary. This enforces a stricter business logic however, thus it is part of the API.
func (db *DB) Put(key string, kind pb.DataType, val proto.Message, meta proto.Message, option *Option) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if option == nil {
		option = defaultOpt()
	}

	return db.put(key, kind, val, meta, option)
}

// put does the actual work of put, but is not threadsafe.
func (db *DB) put(key string, kind pb.DataType, val proto.Message, meta proto.Message, option *Option) error {
	chainHash, err := db.ChainHash(key, option)
	if err != nil {
		if err != ErrKeyNotExist {
			// if the object does not exist, an error should be returned. This is expected if the kind is
			// DataTypeCreated, however if the kind is amend, something is wrong.
			if kind == pb.DataType_Amend {
				return errors.Wrap(err, "unable to find obj, sure it should be amended?")
			}
		}
		// if this is the first obj, we use a hash value of 0
		chainHash = 0
	}

	if kind == pb.DataType_Amend && chainHash == 0 {
		return errors.New("unable to find chain corresponding to key, sure it should be amended?")
	}

	if kind == pb.DataType_Create && chainHash != 0 {
		return errors.New("chain corresponding to key already exists")
	}

	obj, err := ptypes.MarshalAny(val)
	if err != nil {
		return err
	}

	metafields, err := ptypes.MarshalAny(meta)
	if err != nil {
		return err
	}

	object := setObj(key, kind, obj, metafields)
	chunk, err := db.setState(object, chainHash)
	if err != nil {
		return err
	}

	err = db.config.Store.Put(key, chunk)
	if err != nil {
		return err
	}

	db.config.Cache.SetObjHash(key, chunk.State.ObjHash)
	db.config.Cache.SetChainHash(key, chunk.State.ChainHash)

	return nil
}

// makeObj sets the Object keys
func setObj(key string, kind pb.DataType, any *any.Any, meta *any.Any) pb.Object {
	return pb.Object{
		Key:   key,
		Value: any,
		Meta: &pb.Meta{
			Timestamp: ptypes.TimestampNow(),
			Fields:    meta,
			Type:      kind,
		},
	}
}

// setState hashes, embeds and signs the object; returning a writable chunk.
func (db *DB) setState(object pb.Object, prevchainHash uint64) (*pb.Chunk, error) {

	objHash, err := db.config.Hasher.Hash(object)
	if err != nil {
		return nil, err
	}

	chainHash, err := db.config.Hasher.CombineHash(objHash, prevchainHash)
	if err != nil {
		return nil, err
	}

	state := pb.State{
		ObjHash:   objHash,
		ChainHash: chainHash,
		Signature: []byte{},
	}

	chunk := &pb.Chunk{
		Obj:   &object,
		State: &state,
	}

	chunk.State.Signature, err = db.config.Signer.Sign(chunk)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}
