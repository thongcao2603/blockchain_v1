package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thongcao2603/blockchain_v1/crypto"
	"github.com/thongcao2603/blockchain_v1/types"
	"testing"
	"time"
)

func RandomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}

	return NewBlock(header, []Transaction{})
}

func RandomBlockWithSignature(t *testing.T, height uint32, prevBlockhash types.Hash) *Block {
	b := RandomBlock(height, prevBlockhash)
	privateKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privateKey))
	return b
}

func TestHashBlock(t *testing.T) {
	b := RandomBlock(1, types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	b := RandomBlockWithSignature(t, uint32(1), types.Hash{})

	assert.Nil(t, b.Verify())

	//assert.Equal(t, b.Validator, privateKey.PublicKey())
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {

}
