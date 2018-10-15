package install

import (
	"fmt"

	"github.com/carapace/core/internal/scheme"
	"github.com/carapace/core/pkg/handlers/mock"
)

func init() {
	fmt.Println("WARN: mock handler is enabled")
	scheme.Register(mock.Version, mock.Kind, &mock.Handler{}, mock.Validate)
}
