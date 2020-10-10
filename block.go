package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct{
	Timestamp 		int64
	Data 			[]byte
	PrevBlockHash	[]byte
	Hash 			[]byte
}

// *Block <-- 리턴값
// := 변수초기화하면서 선업

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join( [][]byte{b.PrevBlockHash,b.Data, timestamp} ,[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//제너시스 블록 생성
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}