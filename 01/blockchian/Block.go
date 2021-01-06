package blockchian

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	TimeStamp           int64
	Hash                []byte
	PrevBlockHash       []byte
	Height              int64
	Data                []byte
}

func NewBlock(height int64, prevBlockhash []byte, data []byte) *Block {
	var block Block
    block = Block{
    	TimeStamp:time.Now().Unix(),
    	Hash:nil,
    	PrevBlockHash:prevBlockhash,
    	Height:height,
    	Data:data,
	}
    block.SetHash()
	return &block
}

func (b *Block) SetHash() {
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes := IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	}, []byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}

