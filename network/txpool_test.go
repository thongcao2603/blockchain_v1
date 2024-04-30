package network

import (
	"github.com/stretchr/testify/assert"
	"github.com/thongcao2603/blockchain_v1/core"
	"testing"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)
}

func TestTxPool_AddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, p.Add(tx))
}
