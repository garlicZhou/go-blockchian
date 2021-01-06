package blockchian

type BlockChian struct {
	Blocks []*Block
}

func CreateBlockChainWithGeneisiBlock() *BlockChian {
	block := CreateGenesisBlock([]byte("init blockchain"))
	return &BlockChian{Blocks:[]*Block{block}}
}

func (bc *BlockChian) AddBlock(height int64, prevBlockHash []byte, data []byte)  {
	newBlock := NewBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}