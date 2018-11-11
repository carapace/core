package chaindb

import (
	"fmt"
	"testing"

	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDB_Put_Create_NoCache(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	err := db.Put(
		"TestDB_Put_Create_NoCache",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_Put_Create_NoCache"},
		&pb.Test{Value: "META: TestDB_Put_Create_NoCache"},
		&Option{Cached: false},
	)
	require.NoError(t, err)
}

func TestDB_Put_Create_Cached(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	err := db.Put(
		"TestDB_Put_Create_Cached",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_Put_Create_Cached"},
		&pb.Test{Value: "META: TestDB_Put_Create_Cached"},
		&Option{Cached: true},
	)
	require.NoError(t, err)
}

func TestDB_Put_Create_NoOpt(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	err := db.Put(
		"TestDB_Put_Create_NoOpt",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_Put_Create_NoOpt"},
		&pb.Test{Value: "META: TestDB_Put_Create_NoOpt"},
		&Option{Cached: false},
	)
	require.NoError(t, err)
}

func TestDB_Put_Double_Create_Fails_No_Cache(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	err := db.Put(
		"TestDB_Put_Double_Create_Fails_No_Cache",
		pb.DataType_Create, &pb.Test{Value: "TestDB_Put_Double_Create_Fails"},
		&pb.Test{Value: "META: TestDB_Put_Double_Create_Fails"},
		&Option{Cached: false},
	)
	require.NoError(t, errors.Cause(err))

	err = db.Put(
		"TestDB_Put_Double_Create_Fails_No_Cache",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_Put_Double_Create_Fails_2"},
		&pb.Test{Value: "META: TestDB_Put_Double_Create_Fails_2"},
		&Option{Cached: false},
	)
	assert.Error(t, err)
}

func TestDB_Put_Double_Create_Fails_Cached(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	err := db.Put(
		"TestDB_Put_Double_Create_Fails_Cached",
		pb.DataType_Create, &pb.Test{Value: "TestDB_Put_Double_Create_Fails_Cached"},
		&pb.Test{Value: "META: TestDB_Put_Double_Create_Fails_Cached"},
		&Option{Cached: true},
	)
	require.NoError(t, errors.Cause(err))

	err = db.Put(
		"TestDB_Put_Double_Create_Fails_Cached",
		pb.DataType_Create,
		&pb.Test{Value: "TestDB_Put_Double_Create_Fails_Cached_2"},
		&pb.Test{Value: "META: TestDB_Put_Double_Create_Fails_Cached_2"},
		&Option{Cached: true},
	)
	assert.Error(t, err)
}

func TestDB_Put_Create_Then_Amends_Cached(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	tcs := []struct {
		key      string
		dataType pb.DataType
		val      proto.Message
		meta     proto.Message
		option   *Option
	}{
		{
			key:      "TestDB_Put_Create_Then_Amends_Cached",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Cached"},
			option:   &Option{Cached: true},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Cached",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Cached"},
			option:   &Option{Cached: true},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Cached",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Cached"},
			option:   &Option{Cached: true},
		},
	}

	err := db.Put(tcs[0].key, tcs[0].dataType, tcs[0].val, tcs[0].meta, tcs[0].option)
	require.NoError(t, errors.Cause(err))

	for i, tc := range tcs[1:] {
		err := db.Put(tc.key, tc.dataType, tc.val, tc.meta, tc.option)
		assert.NoError(t, err, fmt.Sprintf("error at testcase: %d", i))
	}
}

func TestDB_Put_Create_Then_Amends_No_Cached(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	tcs := []struct {
		key      string
		dataType pb.DataType
		val      proto.Message
		meta     proto.Message
		option   *Option
	}{
		{
			key:      "TestDB_Put_Create_Then_Amends_No_Cache",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_No_Cache"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_No_Cache"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_No_Cache",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_No_Cache"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_No_Cache"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_No_Cache",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_No_Cache"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_No_Cache"},
			option:   &Option{Cached: false},
		},
	}

	err := db.Put(tcs[0].key, tcs[0].dataType, tcs[0].val, tcs[0].meta, tcs[0].option)
	require.NoError(t, errors.Cause(err))

	for i, tc := range tcs[1:] {
		err := db.Put(tc.key, tc.dataType, tc.val, tc.meta, tc.option)
		assert.NoError(t, err, fmt.Sprintf("error at testcase: %d", i))
	}
}

func TestDB_Put_Create_Then_Amends_Mixed_Keys_No_Cache(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()

	tcs := []struct {
		key      string
		dataType pb.DataType
		val      proto.Message
		meta     proto.Message
		option   *Option
	}{
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys2",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys"},
			option:   &Option{Cached: false},
		},
	}

	err := db.Put(tcs[0].key, tcs[0].dataType, tcs[0].val, tcs[0].meta, tcs[0].option)
	require.NoError(t, errors.Cause(err))

	for i, tc := range tcs[1:] {
		err := db.Put(tc.key, tc.dataType, tc.val, tc.meta, tc.option)
		assert.NoError(t, err, fmt.Sprintf("error at testcase: %d", i))
	}
}

func TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached(t *testing.T) {
	db, cleanup := getDB(t)
	defer cleanup()
	tcs := []struct {
		key      string
		dataType pb.DataType
		val      proto.Message
		meta     proto.Message
		option   *Option
	}{
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			option:   &Option{Cached: true},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			option:   &Option{Cached: true},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached2",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			option:   &Option{Cached: true},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			option:   &Option{Cached: true},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_Cached"},
			option:   &Option{Cached: true},
		},
	}

	for i, tc := range tcs {
		err := db.Put(tc.key, tc.dataType, tc.val, tc.meta, tc.option)
		assert.NoError(t, err, fmt.Sprintf("error at testcase: %d", i))
	}
}

func TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache(t *testing.T) {
	tcs := []struct {
		key      string
		dataType pb.DataType
		val      proto.Message
		meta     proto.Message
		option   *Option
	}{
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2",
			dataType: pb.DataType_Create,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2"},
			option:   &Option{Cached: false},
		},
		{
			key:      "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2",
			dataType: pb.DataType_Amend,
			val:      &pb.Test{Value: "TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2"},
			meta:     &pb.Test{Value: "META: TestDB_Put_Create_Then_Amends_Mixed_Keys_With_Closing_No_Cache2"},
			option:   &Option{Cached: false},
		},
	}

	var cleanup func()
	var db *DB
	for i, tc := range tcs {
		db, cleanup = getDB(t)
		err := db.Put(tc.key, tc.dataType, tc.val, tc.meta, tc.option)
		assert.NoError(t, err, fmt.Sprintf("error at testcase: %d", i))
		db.Close()
	}
	cleanup()
}
