package main

import (
	"fmt"
)

type Block struct {
	ID   int
	Data string
}

var blockchain []Block

func DisplayAllBlocks() {
	for _, block := range blockchain {
		fmt.Printf("Block ID: %d, Data: %s\n", block.ID, block.Data)
	}
}

func NewBlock(id int, data string) {
	newBlock := Block{
		ID:   id,
		Data: data,
	}
	blockchain = append(blockchain, newBlock)
}

func ModifyBlock(id int, newData string) {
	for i, block := range blockchain {
		if block.ID == id {
			blockchain[i].Data = newData
			return
		}
	}
	fmt.Println("Block not found")
}

func main() {

	NewBlock(1, "First block data")
	NewBlock(2, "Second block data")

	DisplayAllBlocks()

	ModifyBlock(2, "Modified second block data")
	DisplayAllBlocks()
}
