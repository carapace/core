package append

import (
	"github.com/carapace/cellar"
	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/proto"
)

// ObjectHash returns the latest object hash belonging to a key
func (db *DB) ObjectHash(key string, option *Option) (uint64, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.objectHash(key, option)
}

func (db *DB) objectHash(key string, option *Option) (uint64, error) {
	if option.Cached {
		return db.config.Cache.GetObjHash(key)
	}

	chain, err := db.get(key, option)
	if err != nil {
		return 0, err
	}

	return chain[chain.Len()-1].State.ObjHash, nil
}

// ChainHash returns the current ChainHash belonging to a key
func (db *DB) ChainHash(key string, option *Option) (uint64, error) {
	if option.Cached {
		return db.config.Cache.GetChainHash(key)
	}

	chain, err := db.get(key, option)
	if err != nil {
		return 0, err
	}
	return chain[chain.Len()-1].State.ChainHash, nil

}

func (db *DB) Get(key string, option *Option) (Chain, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.get(key, option)
}

func (db *DB) get(key string, option *Option) (Chain, error) {
	reader := db.cellar.Reader()

	chunks := []*pb.Chunk{}
	err := reader.Scan(func(pos *cellar.ReaderInfo, data []byte) error {
		chunk := &pb.Chunk{}
		err := proto.Unmarshal(data, chunk)
		if err != nil {
			return err
		}

		if chunk.Obj.Key == key {
			chunks = append(chunks, chunk)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(chunks) == 0 {
		return nil, ErrKeyNotExist
	}

	ok, _, err := db.CheckIntegrity(chunks)
	if err != nil {

		if !ok {
			return chunks, err
		}
		// no need to return the chain if the error was not an integrity error
		return nil, err
	}
	return chunks, nil
}
