package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypairSignVerify(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	//address := publicKey.Address()

	msg := []byte("hello world")
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(publicKey, msg))
	fmt.Println(sig)
}
