package blockchian

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/labstack/gommon/log"
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

//序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(block); nil != err {
		log.Panicf("serialize failed %v\n", err)
	}
	return  buffer.Bytes()
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); nil != err {
		log.Panicf("deserialize failed %v\n", err)
	}
	return &block
}


