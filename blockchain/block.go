package blockchain

// BlockChain Struct
type BlockChain struct {
	Blocks []*Block
}

// Block Struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock set up a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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
