package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("block already exists")
	}

	if b.Height != v.bc.Height() {
		return fmt.Errorf("invalid height")
	}

	prevHeader, err := v.bc.GetHeader(b.Height)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)

	if hash != b.PrevBlockHash {
		return fmt.Errorf("invalid prev hash")
	}

	if err := b.Verify(); err != nil {
		return err
	}
	return nil
}
