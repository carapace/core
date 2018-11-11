package cellar

import (
	"github.com/carapace/cellar"
	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/proto"
)

type Engine struct {
	cellar *cellar.DB
}

func (s *Engine) Put(key string, data *pb.Chunk) error {
	ser, err := proto.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.cellar.Append(ser)
	if err != nil {
		return err
	}

	err = s.cellar.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (s *Engine) Get(key string) (chunks []pb.Chunk, err error) {
	reader := s.cellar.Reader()

	err = reader.Scan(func(pos *cellar.ReaderInfo, data []byte) error {
		chunk := &pb.Chunk{}
		err := proto.Unmarshal(data, chunk)
		if err != nil {
			return err
		}

		if chunk.Obj.Key == key {
			chunks = append(chunks, *chunk)
		}
		return nil
	})
	return
}
