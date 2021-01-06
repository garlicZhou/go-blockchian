package main

import "fmt"

func main()  {
	bc := blockchian.CreateBlockChainWithGeneisiBlock()
	fmt.Printf("blockchain: %v\n", bc.Blocks[0])
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1, bc.Blocks[len(bc.Blocks) - 1].Hash, []byte("alice"))
	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1, bc.Blocks[len(bc.Blocks) - 1].Hash, []byte("bob"))

	for _, block := range bc.Blocks {
		fmt.Printf("prevBlockHash: %x, blockHash: %x\n", block.PrevBlockHash, block.Hash)
	}
}
