package core

import "github.com/thongcao2603/blockchain_v1/crypto"

type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}
