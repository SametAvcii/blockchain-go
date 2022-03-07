package main

import (
	"fmt"
	"github.com/SametAvcii/blockchain-go.git/blockchain"
	"os"
	"runtime"
	"strconv"
)

type CommandLine struct {
	blockchain *blockchain.Blockchain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println("print - Prints the blocks in the chain ")

}
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}
func (cli *CommandLine) addBlock(data string) {

}
func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
