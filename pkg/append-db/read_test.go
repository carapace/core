package append

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDB_Read(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)

	msg := &Test{Value: "some random stuff"}
	any, err := ptypes.MarshalAny(msg)
	require.NoError(t, err)

	// generate some test data
	_, err = db.New("TestDB_Read", *any, *any)
	require.NoError(t, err)

	var exists bool
	err = db.Read(func(chunk Chunk) error {
		if chunk.Obj.Key == "TestDB_Read" {
			exists = true
			return nil
		}
		return nil
	})
	require.NoError(t, err)
	assert.True(t, exists)
}

func TestDB_Get(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)

	msg := &Test{Value: "some random stuff"}
	any, err := ptypes.MarshalAny(msg)
	require.NoError(t, err)

	// generate some test data
	_, err = db.Put("TestDB_Get", *any, *any)
	require.NoError(t, err)

	// amend the previous data with a new value
	msg = &Test{Value: "some other stuff"}
	any, err = ptypes.MarshalAny(msg)
	require.NoError(t, err)

	_, err = db.Put("TestDB_Get", *any, *any)
	require.NoError(t, err)

	// Get should return the compacted new chunk, meaning it has the amended object.
	chunk, err := db.Get("TestDB_Get")
	result := &Test{}
	err = ptypes.UnmarshalAny(chunk.Obj.Value, result)
	require.NoError(t, err)

	assert.Equal(t, "some other stuff", result.Value)
}

func TestDB_GetChunks(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)

	msg := &Test{Value: "some random stuff"}
	any, err := ptypes.MarshalAny(msg)
	require.NoError(t, err)

	// generate some test data
	_, err = db.Put("TestDB_GetChunks", *any, *any)
	require.NoError(t, err)

	// amend the previous data with a new value
	msg = &Test{Value: "some other stuff"}
	any, err = ptypes.MarshalAny(msg)
	require.NoError(t, err)

	_, err = db.Put("TestDB_GetChunks", *any, *any)
	require.NoError(t, err)

	// Get should return the compacted new chunk, meaning it has the amended object.
	chunks, err := db.GetChunks("TestDB_GetChunks")
	assert.Equal(t, 2, len(chunks))

	// check if order is preserved
	result := &Test{}
	err = ptypes.UnmarshalAny(chunks[0].Obj.Value, result)
	require.NoError(t, err)
	assert.Equal(t, "some random stuff", result.Value)

	err = ptypes.UnmarshalAny(chunks[1].Obj.Value, result)
	require.NoError(t, err)
	assert.Equal(t, "some other stuff", result.Value)
}
