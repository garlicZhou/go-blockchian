package blockchian

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
)

//db
const dbName = "block.db"
const blockTableName = "blocks"


type BlockChian struct {
	//Blocks []*Block
	DB     *bolt.DB

	Tip    []byte //最新区块的哈希值
}

func CreateBlockChainWithGeneisiBlock() *BlockChian {
    var blockHash []byte
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panicf("create db [%s] failed %v\n", dbName, err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			b, err := tx.CreateBucket([]byte(blockTableName))
			if nil != err {
				log.Panicf("create bucket [%s] failed %v\n", blockTableName, err)
			}

			genesisBlock := CreateGenesisBlock([]byte("init blockchain"))

            err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if nil != err {
				log.Panicf("insert the genesis block failed %v\n", err)
			}
            blockHash = genesisBlock.Hash
			fmt.Printf("blcokh: %x\n", blockHash)
            //存储最新区块哈希
            err = b.Put([]byte("1"), genesisBlock.Hash)
			if nil != err {
				log.Panicf("save the hash of genesis block failed %v\n", err)
			}
		}
		return nil
	})
	return &BlockChian{DB:db, Tip:blockHash}
}

func (bc *BlockChian) AddBlock(data []byte)  {
	bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			hash := b.Get([]byte("1"))
			blockBytes := b.Get(hash)
			latest_block := DeserializeBlock(blockBytes)
			newBlock := NewBlock(latest_block.Height + 1, latest_block.Hash, data)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("insert the newBlock failed %v\n", err)
			}
			err = b.Put([]byte("1"), newBlock.Hash)
			if nil != err {
				log.Panicf("insert the new hash failed %v\n", err)
			}
			bc.Tip = newBlock.Hash
		}
		return nil
	})
}