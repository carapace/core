package append

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLogger(t *testing.T) {
	assert.NotPanics(t, func() {
		defaultLogger()
	})
}
