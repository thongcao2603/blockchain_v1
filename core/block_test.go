package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thongcao2603/blockchain_v1/crypto"
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

func TestSignBlock(t *testing.T) {
	b := RandomBlock(1)
	privateKey := crypto.GeneratePrivateKey()
	assert.Nil(t, b.Sign(privateKey))

	assert.Nil(t, b.Verify())

	assert.Equal(t, b.Validator, privateKey.PublicKey())
	assert.NotNil(t, b.Signature)
}
