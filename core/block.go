package core

import (
	"encoding/binary"
	"io"

	"github.com/thongcao2603/blockchain_v1/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

type Block struct {
	Header
	Transactions []Transaction
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
