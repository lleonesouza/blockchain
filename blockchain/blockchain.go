package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// BlockChain Struct
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// BlockChainIterator Struct S
type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// InitBlockChain with Genesis Block
func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)

	HandleErr(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")

			genesis := Genesis()
			fmt.Println("Genesis proved")

			err = txn.Set(genesis.Hash, genesis.Serialize())
			HandleErr(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			HandleErr(err)
			err = item.Value(func(val []byte) error {
				return err
			})
			return err
		}
	})

	HandleErr(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

// AddBlock add a block to the chain
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		var item, err = txn.Get([]byte("lh"))
		HandleErr(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return err
		})
		return err
	})

	HandleErr(err)
	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		HandleErr(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})

	HandleErr(err)
}

// Iterator go and iterate throw the blockchain and return blocks
func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

// Next call the next block
func (iter *BlockChainIterator) Next() *Block {
	var block *Block
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		HandleErr(err)

		err = item.Value(func(val []byte) error {
			var encodedBlock = val
			block = Deserialize(encodedBlock)
			return err
		})

		return err
	})

	HandleErr(err)

	iter.CurrentHash = block.PrevHash

	return block
}
