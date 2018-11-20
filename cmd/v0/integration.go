package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}
