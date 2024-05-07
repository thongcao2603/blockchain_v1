package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/thongcao2603/blockchain_v1/core"
	"github.com/thongcao2603/blockchain_v1/crypto"
	"github.com/thongcao2603/blockchain_v1/network"
	"log"
	"math/rand"
	"strconv"
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
			if err := sendTransaction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	privKey := crypto.GeneratePrivateKey()
	opts := network.ServerOpts{
		PrivateKey: &privKey,
		ID:         "LOCAL",
		Transports: []network.Transport{trLocal},
	}

	s, err := network.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privateKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(100000000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privateKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())
	return tr.SendMessage(to, msg.Bytes())

}
