package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sign_Verify_Success(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	// address := publicKey.Address()

	msg := []byte("hello world")
	sig, err := privateKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, sig.Verify(publicKey, msg))
	fmt.Println(sig)
}

func TestKeypair_Sign_Verify_Fail(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	// address := publicKey.Address()

	msg := []byte("hello world")
	sig, err := privateKey.Sign(msg)

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	assert.Nil(t, err)
	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("")))
}
