package carapace

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core/auth"
	"hash/fnv"
	"math/rand"
)

type SignOpt struct {
	Recovery bool
}

func (c *Client) SignConfig(config *v0.Config, opt *SignOpt) error {
	if opt == nil {
		opt = &SignOpt{Recovery: false}
	}
	unsigned, err := auth.UnSign(*config)
	if err != nil {
		return err
	}

	sig, err := c.Signer.Sign(c.PrivKey, unsigned)
	if err != nil {
		return err
	}

	key, err := c.PublicKey()
	if err != nil {
		return err
	}

	signature := &v0.Signature{
		S: sig.S,
		R: sig.R,
	}

	if opt.Recovery {
		signature.Key = &v0.Signature_RecoveryPublicKey{RecoveryPublicKey: key}
	} else {
		signature.Key = &v0.Signature_PrimaryPublicKey{PrimaryPublicKey: key}
	}

	if config.Witness == nil {
		config.Witness = &v0.Witness{Signatures: []*v0.Signature{}}
	}

	config.Witness.Signatures = append(config.Witness.Signatures, signature)
	return nil
}

func hash(s string) int64 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	return int64(h.Sum32())
}

func (c *Client) GenPrivKey() error {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.New(rand.NewSource(hash(c.Seed))))
	if err != nil {
		return err
	}
	c.PrivKey = key

	key, err = ecdsa.GenerateKey(elliptic.P256(), rand.New(rand.NewSource(hash(c.RSeed))))
	if err != nil {
		return err
	}
	c.RecoveryPrivKey = key
	return nil
}
