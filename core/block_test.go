package core

import (
	"fmt"
	"github.com/thongcao2603/blockchain_v1/types"
	"testing"
	"time"
)

func RandomBlock(height uint32) *Block {
	header := &Header{
		Height:        height,
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Timestamp:     time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}

func TestHashBlock(t *testing.T) {
	b := RandomBlock(1)
	fmt.Println(b.Hash(BlockHasher{}))
}
