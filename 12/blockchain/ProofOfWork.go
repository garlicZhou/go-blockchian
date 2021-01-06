package blockchian

import (
	"bytes"
	"crypto/sha256"

	"fmt"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block *Block
	//难度哈希
	target *big.Int
}

func NewPOW(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256 - targetBit)
	return &ProofOfWork{Block:block, target:target}
}

//返回哈希值和碰撞次数
func (p *ProofOfWork) Run() ([]byte, int) {
	var nonce = 0
	var hashInt big.Int
	var hash [32]byte
	for {
		dataBytes := p.prepareDate(int64(nonce))
		hash = sha256.Sum256(dataBytes)
        hashInt.SetBytes(hash[:])
		//检测生成的哈希值是否符合条件
		if p.target.Cmp(&hashInt) == 1{
			break
		}
		nonce++
	}
	fmt.Printf("\n碰撞次数:%d\n", nonce)
	return hash[:], nonce
}

func (p *ProofOfWork) prepareDate(nonce int64) []byte {
	timeStampBytes := IntToHex(p.Block.TimeStamp)
	heightBytes := IntToHex(p.Block.Height)
	data := bytes.Join([][]byte{
		timeStampBytes,
		heightBytes,
		p.Block.PrevBlockHash,
		p.Block.HashTransaction(),
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})
	return data
}