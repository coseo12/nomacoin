package blockchain

import (
	"sync"
)

type blockchain struct {
	NewstHash string `json:"newstHash"`
	Height    int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewstHash, b.Height)
	b.NewstHash = block.Hash
	b.Height = block.Height
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
