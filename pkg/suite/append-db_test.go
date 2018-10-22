package test

import (
	"testing"
)

func TestAppendDB(t *testing.T) {
	_, exit := AppendDB()
	defer exit()
}
