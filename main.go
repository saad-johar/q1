package main

import (
	"fmt"
	"time"
)

type Block struct {
	ID        int
	Data      string
	Timestamp time.Time
}

var blockchain []Block

func DisplayAllBlocks() {
	for _, block := range blockchain {
		fmt.Printf("Block ID: %d, Data: %s, Timestamp: %s\n", block.ID, block.Data, block.Timestamp.Format(time.RFC3339))
	}
}

func NewBlock(id int, data string) {
	newBlock := Block{
		ID:        id,
		Data:      data,
		Timestamp: time.Now(),
	}
	blockchain = append(blockchain, newBlock)
}

func ModifyBlock(id int, newData string) error {
	for i, block := range blockchain {
		if block.ID == id {
			blockchain[i].Data = newData
			blockchain[i].Timestamp = time.Now()
			return nil
		}
	}
	return fmt.Errorf("block with ID %d not found", id)
}

func FindBlockByID(id int) (*Block, error) {
	for _, block := range blockchain {
		if block.ID == id {
			return &block, nil
		}
	}
	return nil, fmt.Errorf("block with ID %d not found", id)
}

func main() {

	NewBlock(1, "First block data")
	NewBlock(2, "Second block data")

	DisplayAllBlocks()

	err := ModifyBlock(2, "Modified second block data")
	if err != nil {
		fmt.Println(err)
	}
	DisplayAllBlocks()

	block, err := FindBlockByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found block: ID %d, Data %s\n", block.ID, block.Data)
	}
}
