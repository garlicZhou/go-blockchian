package main

import (
	"go-bc/03/blockchain"
)

func main()  {
	bc := blockchian.CreateBlockChainWithGeneisiBlock()
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1, bc.Blocks[len(bc.Blocks) - 1].Hash, []byte("alice"))
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1, bc.Blocks[len(bc.Blocks) - 1].Hash, []byte("bob"))

}
