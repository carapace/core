package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// Dir creates a directory for the current test. Optionally a name may be passed for the directory
func Dir(t *testing.T, name ...string) (string, func()) {
	if len(name) > 1 {
		panic("provide one name for the directory")
	}

	dirname := "test"

	if name != nil {
		dirname = name[0]
	}

	fullpath := fmt.Sprintf("carapacecoretesting/%s/%s", string(hash(caller())), dirname)
	err := os.MkdirAll(fullpath, 0700)
	require.NoError(t, err)
	return fullpath, func() {
		os.RemoveAll(fullpath)
	}
}
