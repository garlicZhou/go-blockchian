package blockchian

import (
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
)

type BlockChainIterator struct {
	DB             *bolt.DB
	CurrentHash    []byte
}

func (bc *BlockChian) Iterator() *BlockChainIterator {
	return &BlockChainIterator{DB:bc.DB, CurrentHash:bc.Tip}
}

func (bcit *BlockChainIterator) Next() *Block {
	var block *Block

	err := bcit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockBytes := b.Get(bcit.CurrentHash)
			block = DeserializeBlock(currentBlockBytes)
			bcit.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panicf("iterator the db failed %v\n", err)
	}
	return block
}