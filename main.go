package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/thongcao2603/blockchain_v1/core"
	"github.com/thongcao2603/blockchain_v1/crypto"
	"math/rand"
	"strconv"
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
			if err := sendTransaction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
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

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privateKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privateKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}
	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())

	return tr.SendMessage(to, msg.Bytes())
}
