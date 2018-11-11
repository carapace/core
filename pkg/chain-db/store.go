package chaindb

import pb "github.com/carapace/core/pkg/chain-db/proto"

type StorageEngine interface {
	Put(key string, data *pb.Chunk) error
	Get(key string) (chunks []pb.Chunk, err error)
	Close() error
}
