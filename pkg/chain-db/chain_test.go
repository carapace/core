package chaindb

import (
	"fmt"
	"testing"

	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newChain() Chain {
	return Chain{
		{
			State: &pb.State{},
			Obj: &pb.Object{
				Meta: &pb.Meta{
					Timestamp: ptypes.TimestampNow(),
				},
				Key: "key",
			},
		},
		{
			State: &pb.State{},
			Obj: &pb.Object{
				Meta: &pb.Meta{
					Timestamp: ptypes.TimestampNow(),
				},
				Key: "key",
			},
		},
		{
			State: &pb.State{},
			Obj: &pb.Object{
				Meta: &pb.Meta{
					Timestamp: ptypes.TimestampNow(),
				},
				Key: "key",
			},
		},
		{
			State: &pb.State{},
			Obj: &pb.Object{
				Meta: &pb.Meta{
					Timestamp: ptypes.TimestampNow(),
				},
				Key: "key",
			},
		},
	}
}

func genChain(t *testing.T, db *DB, length int) Chain {
	chain := Chain{}

	// firs generate the first one, as it is a bit different
	obj := &pb.Object{Meta: &pb.Meta{Type: pb.DataType_Create, Timestamp: ptypes.TimestampNow()}}
	chunk, err := db.setState(*obj, 0)
	require.NoError(t, err)
	chain = append(chain, chunk)

	for i := 1; i < length-1; i++ {
		obj := &pb.Object{Meta: &pb.Meta{Type: pb.DataType_Create, Timestamp: ptypes.TimestampNow()}}
		chunk, err := db.setState(*obj, chain[i-1].State.ChainHash)
		require.NoError(t, err)

		chain = append(chain, chunk)
	}
	return chain
}

func TestChain_Less(t *testing.T) {
	chain := newChain()
	assert.True(t, chain.Less(1, 2))
}

func TestChain_Len(t *testing.T) {
	chain := newChain()
	assert.True(t, chain.Len() == 4)
}

func TestDB_CheckIntegrity(t *testing.T) {
	db := getDB(t)

	// chain integrity is okay now
	chain := genChain(t, db, 10)
	ok, i, err := db.CheckIntegrity(chain)
	require.NoError(t, err, fmt.Sprintf("incorrect hash at: %d", i))
	assert.Empty(t, i)
	assert.True(t, ok)

	// altering one of the obj hashes
	chain = genChain(t, db, 10)
	chain[3].State.ObjHash = 0
	ok, i, err = db.CheckIntegrity(chain)
	assert.EqualError(t, ErrIncorrectSignature, err.Error())
	assert.Equal(t, 3, i)
	assert.False(t, ok)

	// altering a chain hash
	chain = genChain(t, db, 10)
	chain[3].State.ChainHash = 0
	ok, i, err = db.CheckIntegrity(chain)
	assert.EqualError(t, ErrIncorrectSignature, err.Error())
	assert.Equal(t, 3, i)
	assert.False(t, ok)
}
