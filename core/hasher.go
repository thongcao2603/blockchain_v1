package core

import (
	"crypto/sha256"

	"github.com/thongcao2603/blockchain_v1/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(b *Block) types.Hash {

	h := sha256.Sum256(b.HeaderData())

	return types.Hash(h)
}
