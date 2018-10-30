package ecdsa

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func verifyMySig(pub *ecdsa.PublicKey, digest []byte, sig []byte) bool {
	// https://github.com/gtank/cryptopasta/blob/master/sign.go
	dg := sha1.Sum(digest)

	var esig ecdsaSignature
	esig.R.SetString("89498588918986623250776516710529930937349633484023489594523498325650057801271", 0)
	esig.S.SetString("67852785826834317523806560409094108489491289922250506276160316152060290646810", 0)
	return ecdsa.Verify(pub, dg[:], esig.R, esig.S)
}

func TestHexToPrivateKey(t *testing.T) {
	xHexStr := "4bc55d002653ffdbb53666a2424d0a223117c626b19acef89eefe9b3a6cfd0eb"
	yHexStr := "d8308953748596536b37e4b10ab0d247f6ee50336a1c5f9dc13e3c1bb0435727"
	ePubKey, err := HexToPublicKey(xHexStr, yHexStr)
	require.NoError(t, err)

	sig := "3045022071f06054f450f808aa53294d34f76afd288a23749628cc58add828e8b8f2b742022100f82dcb51cc63b29f4f8b0b838c6546be228ba11a7c23dc102c6d9dcba11a8ff2"
	sigHex, err := hex.DecodeString(sig)
	require.NoError(t, err)

	ok := verifyMySig(ePubKey, []byte("This is string to sign"), sigHex)
	assert.True(t, ok)
}
