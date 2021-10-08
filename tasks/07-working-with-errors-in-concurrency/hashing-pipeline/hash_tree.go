package pipeline

import "errors"

var errNothingToHash = errors.New("nothing to hash")

// CalculateHash реализует хеширование входящих элементов по принципу дерева хешей (Merkle tree).
// Если входящий слайс пуст, то возвращает ошибку errNothingToHash.
func CalculateHash(hh []Hashable) (Hash, error) {
	// Реализуй меня. При желании воспользуйся node.
	return nil, nil
}

// node представляет собой узел хеш-дерева.
type node struct {
	left  Hashable
	right Hashable
}

func (n node) Hash() Hash {
	return newHash(append(n.left.Hash(), n.right.Hash()...))
}
