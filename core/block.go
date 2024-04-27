package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"github.com/thongcao2603/blockchain_v1/types"
	"io"
)

type Header struct {
	Version       uint32
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
	Nonce         uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err1 := binary.Write(w, binary.LittleEndian, &h.PrevBlockHash)
	if err1 != nil {
		return err1
	}
	err2 := binary.Write(w, binary.LittleEndian, &h.Timestamp)
	if err2 != nil {
		return err2
	}
	err3 := binary.Write(w, binary.LittleEndian, &h.Height)
	if err3 != nil {
		return err3
	}
	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	err := binary.Read(r, binary.LittleEndian, &h.Version)
	if err != nil {
		return err
	}
	err1 := binary.Read(r, binary.LittleEndian, &h.PrevBlockHash)
	if err1 != nil {
		return err1
	}
	err2 := binary.Read(r, binary.LittleEndian, &h.Timestamp)
	if err2 != nil {
		return err2
	}
	err3 := binary.Read(r, binary.LittleEndian, &h.Height)
	if err3 != nil {
		return err3
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

type Block struct {
	Header
	Transactions []Transaction

	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	b.Header.EncodeBinary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}
	return b.hash
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil
}
func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}
