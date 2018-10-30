package test

import (
	"testing"
)

func TestDir(t *testing.T) {
	_, exit := Dir(t) // Dir exits the test if errors occur
	defer exit()
}
