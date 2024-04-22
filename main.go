package main

import (
	"time"

	"github.com/thongcao2603/blockchain_v1/network"
)

//server
//transport tcp udp
//block
//tx
//keypair

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	err := trLocal.Connect(trRemote)
	if err != nil {
		return
	}
	err1 := trRemote.Connect(trLocal)
	if err1 != nil {
		return
	}

	go func() {
		for {
			err := trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
