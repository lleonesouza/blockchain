package main

import (
	"fmt"
	"strconv"

	"github.com/lleonesouza/blockchain/blockchain"
)

func main() {

	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Secibd Block after Genesis")
	fmt.Println()
	fmt.Println()

	iter := chain.Iterator()

	for {
		block := iter.Next()
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println("")

		if len(block.PrevHash) == 0 {
			break
		}
	}

}
