package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/thongcao2603/blockchain_v1/crypto"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	msg := []byte("foo")
	tx := &Transaction{
		Data: msg,
	}
	privateKey := crypto.GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	assert.Nil(t, tx.Sign(privateKey))
	assert.Equal(t, tx.PublicKey, publicKey)
	assert.Equal(t, tx.Data, msg)
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	msg := []byte("foo")
	tx := &Transaction{
		Data: msg,
	}
	privateKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privateKey))
	assert.Nil(t, tx.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivateKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}
