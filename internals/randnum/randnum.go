package randnum

import (
	"crypto/rand"
	"math/big"
)

// InRange returns random int64 for given range
func InRange(from, to int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(to))
	if err != nil {
		return from
	}

	return nBig.Int64() + from
}
