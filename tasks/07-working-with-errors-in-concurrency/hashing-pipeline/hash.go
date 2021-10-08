package pipeline

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hashable interface {
	Hash() Hash
}

type Hash []byte

func (h Hash) String() string {
	return hex.EncodeToString(h)
}

func newHash(data []byte) Hash {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}
