package main

import (
	"fmt"
	"go-bc/01/blockchian"
)

func main()  {
	block := blockchian.NewBlock(1, nil, []byte("first block"))
	fmt.Printf("block:%x\n",block.Hash)

}