package scheme

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

// Validator defines a function type which checks the body of a configuration message,
// checking for incorrect keys, formatting etc. It should throw an error if the contents of
// the file are incorrect.
type Validator func(v1.Config) error
