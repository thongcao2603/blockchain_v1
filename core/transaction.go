package core

import (
	"fmt"
	"github.com/thongcao2603/blockchain_v1/crypto"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.From = privateKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("signature is nil")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("signature is invalid")
	}

	return nil
}
