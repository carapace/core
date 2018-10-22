package state

import (
	"crypto/hmac"
	"crypto/sha256"
)

// Signer defines a common interface for cryptographically signing objects. Signer is required to be cryptographically
// safe
type Signer interface {
	Sign(interface{}) ([]byte, error)
}

// Verifier defines a common interface for checking cryptographically signed objects.
type Verifier interface {
	Verify(interface{}, []byte) (bool, error)
}

// HMACSigner is both a Signer and Verifier, using hmac(sha256) encoding to sign and verify signatures.
type HMACSigner struct {
	encoder Encoder
	key     []byte
}

// NewHMAC is the constructor for the HMACSigner. By default it uses the JSON encoder
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

// Sign creates a cryptographic signature of the object
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

// Verify checks a cryptographic signature using the orginal object and the provided signature.
//
// It only returns an error on write errors and encoding errors, the boolean indicates if signatures match.
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
