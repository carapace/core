package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedis(t *testing.T) {
	r, exit := Redis()
	defer exit()
	assert.NoError(t, r.Ping().Err())
}
