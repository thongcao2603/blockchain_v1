package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/thongcao2603/blockchain_v1/types"
	"testing"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(RandomBlock(0, types.Hash{}))
	assert.Nil(t, err)
	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	for i := 0; i < 1000; i++ {
		prevHash := getPrevBlockHash(t, bc, uint32(i+1))
		block := RandomBlockWithSignature(t, uint32(i+1), prevHash)
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(1000))

}

func TestBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

}

func TestHashBlock1(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}

//sai cho nay
