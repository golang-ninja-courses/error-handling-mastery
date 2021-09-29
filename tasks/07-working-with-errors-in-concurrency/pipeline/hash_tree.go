package pipeline

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type Hash []byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func hash(data []byte) Hash {
	h := sha256.Sum256(data)
	return h[:]
}

type Hashable interface {
	Hash() Hash
}

type node struct {
	left  Hashable
	right Hashable
}

func (n node) Hash() Hash {
	return hash(append(n.left.Hash(), n.right.Hash()...))
}

var (
	errNothingToHash = errors.New("nothing to hash")
)

// CalculateHash реализует хеширование входящих элементов по принципу хэш-дерева.
// Принцип хеширования визуализирован в поясняющей диаграмме в шаге задачи.
func CalculateHash(hh []Hashable) (Hash, error) {
	// Реализовать.
	return nil, nil
}
