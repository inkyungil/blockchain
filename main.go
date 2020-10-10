package main

import "fmt"

func main()  {
	bc := NewBlockchain()

	bc.AddBlock("send 1 Token to KKK")
	bc.AddBlock("send 1 Token to TEMP")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
