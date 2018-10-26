package chaindb

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDB_ObjectHash_Not_Set_Cached(t *testing.T) {
	db := getDB(t)
	_, err := db.ObjectHash("TestDB_ObjectHash_Not_Set_Cached", &Option{Cached: true})
	assert.Error(t, err)
}

func TestDB_ObjectHash_Not_Set_No_Cached(t *testing.T) {
	db := getDB(t)
	_, err := db.ObjectHash("TestDB_ObjectHash_Not_Set_No_Cached", &Option{Cached: false})
	assert.Error(t, err)
}

func TestDB_ObjectHash_Set_Cached(t *testing.T) {
	db := getDB(t)

	err := db.Put(
		"TestDB_ObjectHash_Set_Cached",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_ObjectHash_Set_Cached"},
		&pb.Test{Value: "META: TestDB_ObjectHash_Set_Cached"},
		&Option{Cached: true},
	)
	require.NoError(t, err)

	_, err = db.ObjectHash("TestDB_ObjectHash_Set_Cached", &Option{Cached: true})
	assert.NoError(t, err)
}

func TestDB_ObjectHash_Set_No_Cache(t *testing.T) {
	db := getDB(t)

	err := db.Put(
		"TestDB_ObjectHash_Set_No_Cache",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_ObjectHash_Set_No_Cache"},
		&pb.Test{Value: "META: TestDB_ObjectHash_Set_No_Cache"},
		&Option{Cached: false},
	)
	require.NoError(t, err)

	_, err = db.ObjectHash("TestDB_ObjectHash_Set_No_Cache", &Option{Cached: true})
	assert.NoError(t, err)
}

func TestDB_ChainHash_Set_Cached(t *testing.T) {
	db := getDB(t)

	err := db.Put(
		"TestDB_ChainHash_Set_Cached",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_ChainHash_Set_Cached"},
		&pb.Test{Value: "META: TestDB_ChainHash_Set_Cached"},
		&Option{Cached: true},
	)
	require.NoError(t, err)

	_, err = db.ChainHash("TestDB_ChainHash_Set_Cached", &Option{Cached: true})
	assert.NoError(t, err)
}

func TestDB_ChainHash_Set_No_Cache(t *testing.T) {
	db := getDB(t)

	err := db.Put(
		"TestDB_ChainHash_Set_No_Cache",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_ChainHash_Set_No_Cache"},
		&pb.Test{Value: "META: TestDB_ChainHash_Set_No_Cache"},
		&Option{Cached: false},
	)
	require.NoError(t, err)

	_, err = db.ChainHash("TestDB_ChainHash_Set_No_Cache", &Option{Cached: true})
	assert.NoError(t, err)
}

func TestDB_Write_Read_No_Cache_Two_Keys(t *testing.T) {
	db := getDB(t)

	tcs := []struct {
		key      string
		dataType pb.DataType
		val      proto.Message
		meta     proto.Message
		option   *Option
	}{
		{
			key:      "TestDB_Write_Read_No_Cache_Two_Keys",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Write_Read_No_Cache_Two_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Write_Read_No_Cache_Two_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Write_Read_No_Cache_Two_Keys",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Write_Read_No_Cache_Two_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Write_Read_No_Cache_Two_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Write_Read_No_Cache_Two_Keys",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Write_Read_No_Cache_Two_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Write_Read_No_Cache_Two_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Write_Read_No_Cache_Two_Keys2",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Write_Read_No_Cache_Two_Keys2"},
			meta:     &pb.Test{Value: "META: TestDB_Write_Read_No_Cache_Two_Keys2"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Write_Read_No_Cache_Two_Keys2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Write_Read_No_Cache_Two_Keys2"},
			meta:     &pb.Test{Value: "META: TestDB_Write_Read_No_Cache_Two_Keys2"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Write_Read_No_Cache_Two_Keys2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Write_Read_No_Cache_Two_Keys2"},
			meta:     &pb.Test{Value: "META: TestDB_Write_Read_No_Cache_Two_Keys2"},
			option:   &Option{Cached: false},
		},
	}

	for i, tc := range tcs {
		err := db.Put(tc.key, tc.dataType, tc.val, tc.meta, tc.option)
		assert.NoError(t, err, fmt.Sprintf("error at testcase: %d", i))
	}

	chain, err := db.Get("TestDB_Write_Read_No_Cache_Two_Keys", &Option{Cached: false})
	require.NoError(t, err)
	for _, c := range chain {
		assert.Equal(t, c.Obj.Key, "TestDB_Write_Read_No_Cache_Two_Keys")
		val := &pb.Test{}
		err = ptypes.UnmarshalAny(c.Obj.Value, val)
		require.NoError(t, err)
		assert.Equal(t, "TestDB_Write_Read_No_Cache_Two_Keys", val.Value)
	}

	chain, err = db.Get("TestDB_Write_Read_No_Cache_Two_Keys2", &Option{Cached: false})
	require.NoError(t, err)
	for _, c := range chain {
		assert.Equal(t, c.Obj.Key, "TestDB_Write_Read_No_Cache_Two_Keys2")
		val := &pb.Test{}
		err = ptypes.UnmarshalAny(c.Obj.Value, val)
		require.NoError(t, err)
		assert.Equal(t, "TestDB_Write_Read_No_Cache_Two_Keys2", val.Value)
	}
}
