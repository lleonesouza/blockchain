package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain Struct
type BlockChain struct {
	Blocks []*Block
}

// Block Struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash generate hash for the block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock set up a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock add a block to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis create a Genesis Block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})

}

// Init the BlockChain
func Init() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}

}
