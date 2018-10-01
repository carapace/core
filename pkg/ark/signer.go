package ark

import (
	ArkV2 "github.com/ArkEcosystem/go-crypto/crypto"
)

func (s *Service) sign(tx *ArkV2.Transaction) {
	tx.Sign(s.FirstSecret)
	if s.SecondSecret != "" {
		tx.SecondSign(s.SecondSecret)
	}
}

func (s *Service) verify(tx *ArkV2.Transaction) (bool, error) {
	return tx.Verify()
}
