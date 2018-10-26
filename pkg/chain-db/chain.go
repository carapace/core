package append

import (
	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Chain []*pb.Chunk

func (c Chain) Len() int { return len(c) }
func (c Chain) Less(i, j int) bool {
	its, err := ptypes.Timestamp(c[i].Obj.Meta.Timestamp)
	if err != nil {
		// think of something to do
		panic(err)
	}
	jts, err := ptypes.Timestamp(c[j].Obj.Meta.Timestamp)
	if err != nil {
		// think of something to do
		panic(err)
	}
	return its.Before(jts)
}
func (c Chain) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

var (
	ErrIncorrectSignature = errors.New("incorrect signature")
	ErrIncorrectHash      = errors.New("incorrect hash")
)

// CheckIntegrity verifies the integrity of an object chain
//
// It does not presort by timestamp (which should never be needed unless chain order
// has been edited manually.
// If the chain integrity is off, the second argument returned points to the index where the
// hashes did not match up.
func (db *DB) CheckIntegrity(chain Chain) (bool, int, error) {
	var prevHash uint64 = 0

	for i, c := range chain {
		// first we check the signature
		ok, err := verifySignature(db, *c, c.State.Signature)
		if err != nil {
			return false, i, err
		}

		if !ok {
			return false, i, ErrIncorrectSignature
		}

		wantedChainHash, err := db.config.Hasher.CombineHash(c.State.ObjHash, prevHash)
		if err != nil {
			return false, i, err
		}

		if c.State.ChainHash == wantedChainHash {
			prevHash = wantedChainHash
			continue
		}
		db.Logger().Error("", zap.Error(err), zap.Uint64("WANT", wantedChainHash), zap.Uint64("GOT", c.State.ChainHash))
		return false, i, ErrIncorrectHash
	}
	return true, 0, nil
}

func verifySignature(db *DB, chunk pb.Chunk, sig []byte) (bool, error) {
	tmp := chunk.State.Signature
	chunk.State.Signature = []byte{}
	ok, err := db.config.Verifier.Verify(chunk, sig)
	chunk.State.Signature = tmp
	return ok, err
}
