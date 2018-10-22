package state

import (
	"crypto/hmac"
	"crypto/sha256"
)

type Signer interface {
	Sign(interface{}) ([]byte, error)
}

type Verifier interface {
	Verify(interface{}, []byte) (bool, error)
}

type HMACSigner struct {
	encoder Encoder
	key     []byte
}

func NewHMAC(secret []byte, opts ...func(signer *HMACSigner) error) *HMACSigner {
	h := &HMACSigner{
		key:     secret,
		encoder: JSONEncoder{},
	}

	for _, opt := range opts {
		err := opt(h)
		if err != nil {
			panic(err)
		}
	}
	return h
}

func (h *HMACSigner) Sign(obj interface{}) ([]byte, error) {
	btes, err := h.encoder.Encode(obj)
	if err != nil {
		return nil, err
	}

	hsh := hmac.New(sha256.New, h.key)
	_, err = hsh.Write(btes)
	if err != nil {
		return nil, err
	}
	return hsh.Sum(nil), nil
}

func (h *HMACSigner) Verify(obj interface{}, signature []byte) (bool, error) {
	btes, err := h.encoder.Encode(obj)
	if err != nil {
		return false, err
	}

	hsh := hmac.New(sha256.New, h.key)
	_, err = hsh.Write(btes)
	if err != nil {
		return false, err
	}
	return hmac.Equal(signature, hsh.Sum(nil)), nil
}
