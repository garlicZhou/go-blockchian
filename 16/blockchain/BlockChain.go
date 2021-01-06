package blockchian

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
	"math/big"
	"os"
	"strconv"
)

//db
const dbName = "block.db"
const blockTableName = "blocks"


type BlockChian struct {
	//Blocks []*Block
	DB     *bolt.DB

	Tip    []byte //最新区块的哈希值
}

func dbExist() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateBlockChainWithGeneisiBlock(address string) *BlockChian {
    if dbExist() {
    	fmt.Println("genesis block is existed")
    	os.Exit(1)
	}
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

			txCoinbase := NewCoinbaseTransaction(address)


			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})

            err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if nil != err {
				log.Panicf("insert the genesis block failed %v\n", err)
			}
            blockHash = genesisBlock.Hash
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

func (bc *BlockChian) AddBlock(txs []*Transaction)  {
	bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			hash := b.Get([]byte("1"))
			blockBytes := b.Get(hash)
			latest_block := DeserializeBlock(blockBytes)
			newBlock := NewBlock(latest_block.Height + 1, latest_block.Hash, txs)
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

func(bc *BlockChian) PrintChain()  {
	fmt.Println("blockchian  total information")
	var curBlock *Block
	bcit := bc.Iterator()
	for {
		fmt.Println("------------------------------")
		curBlock = bcit.Next()
		fmt.Printf("\tHash:%x\n", curBlock.Hash)
		fmt.Printf("\tPrevBlockHash:%x\n", curBlock.PrevBlockHash)
		fmt.Printf("\tTimeStamp:%v\n", curBlock.TimeStamp)
		fmt.Printf("\tHeight:%d\n", curBlock.Height)
		fmt.Printf("\tNonce:%x\n", curBlock.Nonce)
		fmt.Printf("\tTxs:%v\n", curBlock.Txs)
		for _, tx := range curBlock.Txs {
			fmt.Printf("\t\ttx-hash: %x\n", tx.TxHash)
			fmt.Printf("\t\tinput...\n")
			for _, vin := range tx.Vins {
				fmt.Printf("\t\t\tvin-txHash : %x\n", vin.TxHash)
				fmt.Printf("\t\t\tvin-vout : %v\n", vin.Vout)
				fmt.Printf("\t\t\tvin-scriptSig : %s\n", vin.ScriptSig)
			}
			fmt.Printf("\t\toutput...\n")
			for _, vout := range tx.Vouts {
				fmt.Printf("\t\t\tvout-value: %d\n", vout.Value)
				fmt.Printf("\t\t\tvout-scriptPubkey : %s\n", vout.ScriptPubkey)
			}
		}

		var hashInt big.Int
		hashInt.SetBytes(curBlock.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

func BlockchainObject() *BlockChian {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("open db failed %v\n", err)
	}
	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("1"))
		}
		return nil
	})
	if err != nil {
		log.Panicf("set blockchain object failed %v\n", err)
	}
	return &BlockChian{DB:db, Tip:tip}
}

func (bc *BlockChian) MineNewBlock(from, to, amount []string)  {
	var txs []*Transaction
	var block *Block
	//生成新交易

	value, _ := strconv.Atoi(amount[0])
    tx := NewSimpleTransaction(from[0], to[0], value)
    txs = append(txs, tx)

	bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			hash := b.Get([]byte("1"))
			blockBytes := b.Get(hash)
			block = DeserializeBlock(blockBytes)
		}
		return nil
	})

	block = NewBlock(block.Height + 1, block.Hash, txs)

	bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			err := b.Put(block.Hash, block.Serialize())
			if err != nil {
				log.Panicf("add block to db failed %v\n", err)
			}
			err = b.Put([]byte("1"), block.Hash)
			if err != nil {
				log.Panicf("update the latest block hash failed %v\n", err)
			}
			bc.Tip = block.Hash
		}
		return nil
	})
}