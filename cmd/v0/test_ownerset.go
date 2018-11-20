package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"path"
	"time"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s Suite) TestInfoService() {
	core, cleanup := s.NewCore(s.T(), "./TestOwnerSetInfo")
	defer cleanup()

	ctx, cf := context.WithTimeout(context.Background(), 50*time.Millisecond)
	cf()
	info, err := core.InfoService(ctx, &empty.Empty{})
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), info)
	spew.Dump(info)
}

func (s Suite) TestOwnerSetFirstOwnersSimple() {
	core, cleanup := s.NewCore(s.T(), "./TestOwnerSetFirstOwnersSimple")
	defer cleanup()

	ctx, cf := context.WithTimeout(context.Background(), 50*time.Millisecond)
	cf()

	client := s.NewClient()
	err := client.GenPrivKey()
	s.Require().NoError(err)

	ownerset, err := client.LoadConfig(path.Join("testdata", "ownerSet.yaml"))
	s.Require().NoError(err)

	err = client.FillUserKeys(ownerset)
	s.Require().NoError(err)

	err = client.SignConfig(ownerset, nil)
	s.Require().NoError(err)

	// Configuring the node owners for the first time should not return an error
	res, err := core.ConfigService(ctx, ownerset)
	s.Require().NoError(err)
	s.Assert().NotNil(res)
	s.Assert().Equal(v0.Code_OK, res.Code, fmt.Sprintf("MSG: %s, ERR: %s", res.MSG, res.Err))

	// altering the set with an authorized user should correctly alter the ownerSet
	ownerset.Header.Increment += 1
	res, err = core.ConfigService(ctx, ownerset)
	s.Require().NoError(err)
	s.Assert().Equal(v0.Code_OK, res.Code, fmt.Sprintf("MSG: %s, ERR: %s", res.MSG, res.Err))

	// altering the set with an incorrect sig should error
	ownerset.Header.Increment += 1
	ownerset.Witness.Signatures[0].R = []byte("bad signature")
	res, err = core.ConfigService(ctx, ownerset)
	s.Require().NoError(err)
	s.Assert().Equal(v0.Code_BadRequest, res.Code, fmt.Sprintf("MSG: %s, ERR: %s", res.MSG, res.Err))
}
