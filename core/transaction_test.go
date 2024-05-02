package core

import (
	"bytes"
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
	assert.Equal(t, tx.From, publicKey)
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
	tx.From = otherPrivateKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	tx := &Transaction{
		Data: []byte("foo"),
	}
	privateKey := crypto.GeneratePrivateKey()

	assert.Nil(t, tx.Sign(privateKey))
	return tx

}

func TestEncodeDecodeTransaction(t *testing.T) {
	tx := randomTxWithSignature(t)
	buf := &bytes.Buffer{}
	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))

	//tr := new(Transaction)
	//assert.Nil(t, tr.Decode(NewGobTxDecoder(buf)))
	//assert.Equal(t, tr, tx)
}
