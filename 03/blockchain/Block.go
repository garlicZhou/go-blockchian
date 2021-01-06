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
	Nonce               int64
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
    pow := NewPOW(&block)
    hash, nouce := pow.Run()
    block.Hash = hash
    block.Nonce = int64(nouce)
	return &block
}

func (b *Block) SetHash() {
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes := IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		timeStampBytes,
		heightBytes,
		b.PrevBlockHash,
		b.Data,
	}, []byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}

func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}

