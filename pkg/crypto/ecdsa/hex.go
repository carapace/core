package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
)

func HexToPrivateKey(hexStr string) (*ecdsa.PrivateKey, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	k := new(big.Int)
	k.SetBytes(bytes)

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = elliptic.P256()
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(k.Bytes())

	return priv, nil
}

func HexToPublicKey(xHex string, yHex string) (*ecdsa.PublicKey, error) {
	xBytes, err := hex.DecodeString(xHex)
	if err != nil {
		return nil, err
	}

	x := new(big.Int)
	x.SetBytes(xBytes)

	yBytes, err := hex.DecodeString(yHex)
	if err != nil {
		return nil, err
	}

	y := new(big.Int)
	y.SetBytes(yBytes)

	pub := new(ecdsa.PublicKey)
	pub.X = x
	pub.Y = y

	pub.Curve = elliptic.P256()

	return pub, nil
}
