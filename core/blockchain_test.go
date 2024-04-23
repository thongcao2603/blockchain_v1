package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newBlockChainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < 1000; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))

	}
	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)

	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}

func TestBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}
