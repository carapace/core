package append

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDB_New(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)

	msg := &Test{Value: "some random stuff"}
	any, err := ptypes.MarshalAny(msg)
	require.NoError(t, err)

	c, err := db.New("TestDB_New", *any, *any)
	require.NoError(t, err)

	assert.Equal(t, DataType_Create, c.Obj.Meta.Type)

	// Assert we get an error when trying again, because the key now already exists
	_, err = db.New("TestDB_New", *any, *any)
	assert.EqualError(t, err, KeyExists.Error())
}

func TestDB_Put(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)

	msg := &Test{Value: "different random stuff"}
	any, err := ptypes.MarshalAny(msg)
	require.NoError(t, err)

	c, err := db.Put("TestDB_Put", *any, *any)
	require.NoError(t, err)
	assert.Equal(t, DataType_Create, c.Obj.Meta.Type)

	// Put should automatically create an amend if the key already exists
	c, err = db.Put("TestDB_Put", *any, *any)
	require.NoError(t, err)
	assert.Equal(t, DataType_Amend, c.Obj.Meta.Type)
}

func TestDB_Amend(t *testing.T) {
	db, err := New("testdb")
	require.NoError(t, err)

	msg := &Test{Value: "different random stuff"}
	any, err := ptypes.MarshalAny(msg)
	require.NoError(t, err)

	// Key does not exist, thus an error should be thrown
	_, err = db.Amend("TestDB_Amend", *any, *any)
	require.Error(t, err)

	// Put should automatically create an amend if the key already exists
	c, err := db.Put("TestDB_Amend", *any, *any)
	require.NoError(t, err)
	assert.Equal(t, DataType_Create, c.Obj.Meta.Type)

	// Now the key has been created, thus the amend should not error.
	c, err = db.Amend("TestDB_Amend", *any, *any)
	require.NoError(t, err)
	assert.Equal(t, DataType_Amend, c.Obj.Meta.Type)
}
