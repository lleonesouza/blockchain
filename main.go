package main

import (
	"fmt"

	"github.com/lleonesouza/blockchain/blockchain"
)

func main() {
	chain := blockchain.Init()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Secibd Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("")
	}

}
