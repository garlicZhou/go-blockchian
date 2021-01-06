package blockchian

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

type CLI struct {
}

func PrintUsage()  {
	fmt.Println("Usage:")
	fmt.Println("createBlockchain")
	fmt.Println("addBlock -data DATA")
	fmt.Println("printChain")
}

func (cli *CLI) createBlockchain()  {
	CreateBlockChainWithGeneisiBlock()
}

func (cli *CLI) addBlock(data string) {

	if !dbExist() {
		fmt.Println("database is not existed")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	blockchain.AddBlock([]byte(data))
}

func (cli *CLI) printChian()  {
	if !dbExist() {
		fmt.Println("database is not existed")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	blockchain.PrintChain()
}

//检测参数数量

func IsValidArgs()  {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run()  {
	IsValidArgs()
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockchainCmd :=flag.NewFlagSet("createBlockchain", flag.ExitOnError)

	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "add block info")

	switch os.Args[1] {
	case "addBlock":
		if err := addBlockCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("prase addBlockCmd failed! %v\n ", err)
		}

	case "printChain":
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("prase printChainCmd failed! %v\n ", err)
		}

	case "createBlockchain":
		if err := createBlockchainCmd.Parse(os.Args[2:]); err != nil {
			log.Panicf("prase createBlockchainCmd failed! %v\n ", err)
		}

	default:
		PrintUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockArg)
	}

	if printChainCmd.Parsed() {
		cli.printChian()
	}

	if createBlockchainCmd.Parsed() {
        cli.createBlockchain()
	}
}