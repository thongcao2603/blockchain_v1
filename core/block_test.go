package core

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/thongcao2603/blockchain_v1/types"
	"testing"
	"time"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        10,
		Nonce:         123123,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h, hDecode)
}
