package main

import (
	"context"
	"github.com/carapace/core/api/v0/proto"
	"path"
	"time"
)

func (s Suite) TestUserSetRobotSet() {
	core, cleanup := s.NewCore(s.T(), "./TestUserSetRobotSet")
	defer cleanup()

	ctx, cf := context.WithTimeout(context.Background(), 50*time.Millisecond)
	cf()

	client := s.NewClient()
	err := client.GenPrivKey()
	s.Require().NoError(err)

	userSet, err := client.LoadConfig(path.Join("testdata", "userSet.yaml"))
	s.Require().NoError(err)
	err = client.SignConfig(userSet, nil)
	s.Require().NoError(err)

	// first we need to post an ownerset
	ownerSet, err := client.LoadConfig(path.Join("testdata", "ownerSet.yaml"))
	s.Require().NoError(err)

	err = client.FillUserKeys(ownerSet)
	s.Require().NoError(err)

	err = client.SignConfig(ownerSet, nil)
	s.Require().NoError(err)

	res, err := core.ConfigService(ctx, ownerSet)
	s.Require().NoError(err)
	s.Assert().Equal(res.Code, v0.Code_OK)

	// now to actually posting the new users.
	res, err = core.ConfigService(ctx, userSet)
	s.Require().NoError(err)
	s.Assert().Equal(res.Code, v0.Code_OK, res.MSG, res.Err)
}
