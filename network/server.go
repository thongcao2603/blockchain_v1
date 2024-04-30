package network

import (
	"fmt"
	"github.com/thongcao2603/blockchain_v1/crypto"
	"time"
)

type ServerOpts struct {
	Transports []Transport
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts:  opts,
		isValidator: opts.PrivateKey != nil,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}
}
func (s *Server) Start() error {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do something")
		}
	}

	return nil
}

func (s *Server) initTransports() error {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
	return nil
}
