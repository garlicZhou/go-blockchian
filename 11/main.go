package main

import "go-bc/11/blockchain"

func main()  {
	/*bc := blockchian.CreateBlockChainWithGeneisiBlock()
	bc.AddBlock([]byte("A send 100 btc to b"))
	bc.AddBlock([]byte("B send 200 btc to c"))
	bc.PrintChain()*/
	cli := blockchian.CLI{}
	cli.Run()
}
