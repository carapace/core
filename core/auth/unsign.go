package auth

import (
	"github.com/carapace/core/api/v0/proto"
	"github.com/jinzhu/copier"
)

// UnSign returns an unsigned config.
//
// It does not modify the original config
func UnSign(config v0.Config) (*v0.Config, error) {
	var unsigned = &v0.Config{}
	err := copier.Copy(&unsigned, &config)
	if err != nil {
		return nil, err
	}
	emptyWitness := &v0.Witness{}
	emptyWitness.Reset()
	unsigned.Witness = emptyWitness
	return unsigned, nil
}
