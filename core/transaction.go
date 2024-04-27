package core

import "github.com/thongcao2603/blockchain_v1/crypto"

type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PublicKey = privateKey.PublicKey()
	tx.Signature = sig

	return nil
}
