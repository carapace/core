package signing

import (
	"fmt"

	"github.com/carapace/core/api/v1/proto/generated"
)

type AssetService interface {
	SigningService
	ValidationService
}

type Service struct {
	services map[v1.Asset]AssetService
}

// Register adds an AssetService to the registered assets. If the AssetService has
// already been registered, Register will override the previous service, and emit a warning.
func (s *Service) Register(asset v1.Asset, service AssetService) {
	if _, ok := s.services[asset]; ok {
		panic(fmt.Sprintf(`an assetservice has already been registerd for: %s`, asset.String()))
	}
	s.services[asset] = service
}
