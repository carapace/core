package bitcoin

import (
	"io"
)

func (s Service) Sign(reader io.Reader) (io.Reader, error) {

