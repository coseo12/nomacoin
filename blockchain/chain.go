package blockchain

import (
	"sync"

	"github.com/coseo12/nomacoin/db"
	"github.com/coseo12/nomacoin/utils"
)

type blockchain struct {
	NewstHash string `json:"newstHash"`
	Height    int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewstHash, b.Height+1)
	b.NewstHash = block.Hash
	b.Height = block.Height
	b.persist()
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
