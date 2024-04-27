package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("LOCAL")
	trb := NewLocalTransport("REMOTE")

	assert.Nil(t, trb.Connect(tra))

}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("LOCAL")
	trb := NewLocalTransport("REMOTE")

	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, trb.SendMessage(tra.Addr(), msg))

	rpc := <-tra.Consume()

	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, trb.Addr())
}
