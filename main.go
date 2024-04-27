package main

import (
	"github.com/thongcao2603/blockchain_v1/network"
	"time"
)

//server
//transport tcp udp
//block
//tx
//keypair

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trRemote.Connect(trLocal)
	go func() {
		for {
			msg := []byte("hello world")
			trRemote.SendMessage(trLocal.Addr(), msg)
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
