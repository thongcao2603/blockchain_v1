package network

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

func TestBroadcast(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	trc := NewLocalTransport("C")

	tra.Connect(trb)
	tra.Connect(trc)

	msg := []byte("hello world")
	assert.Nil(t, tra.Broadcast(msg))

	rpcb := <-trb.Consume()
	b, err := ioutil.ReadAll(rpcb.Payload)
	assert.Nil(t, err)
	assert.Equal(t, b, msg)

	rpcc := <-trc.Consume()
	c, err1 := ioutil.ReadAll(rpcc.Payload)
	assert.Nil(t, err1)
	assert.Equal(t, c, msg)

}
