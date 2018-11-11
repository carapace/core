package sqlite3

import (
	"testing"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/carapace/core/pkg/suite"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEngine_Migrate(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	e := New(db)

	assert.NoError(t, e.Migrate())
}

func TestEngine_Put(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	e := New(db)

	assert.NoError(t, e.Migrate())
	assert.NoError(t, e.Put("key", &pb.Chunk{}))
}

func mockData(t *testing.T) *any.Any {
	p := pb.Test{Value: "TEST"}
	any, err := ptypes.MarshalAny(&p)
	require.NoError(t, err)

	return any
}

func TestEngine_Put_Get(t *testing.T) {
	db, cleanup := test.Sqlite3(t)
	defer cleanup()

	e := New(db)
	assert.NoError(t, e.Migrate())

	tcs := []struct {
		key  string
		data *pb.Chunk

		desc string
	}{
		{
			key:  "shortkey",
			data: &pb.Chunk{Obj: &pb.Object{Key: "shortkey"}},
			desc: "empty data chunk should work",
		},
		{
			key:  "shortkey2",
			data: &pb.Chunk{Obj: &pb.Object{Key: "shortkey2", Value: mockData(t)}},
			desc: "mock data chunk should work",
		},
	}

	for _, tc := range tcs {
		require.NoError(t, e.Put(tc.key, tc.data))

		chain, err := e.Get(tc.key)
		require.NoError(t, err)
		assert.Equal(t, tc.key, chain[0].Obj.Key)
	}
}
