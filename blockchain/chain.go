package blockchain

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/coseo12/nomacoin/db"
	"github.com/coseo12/nomacoin/utils"
)

type blockchain struct {
	NewestHash        string `json:"NewestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
	m                 sync.Mutex
}

type storage interface {
	FindBlock(hash string) []byte
	LoadChain() []byte
	SaveBlock(hash string, data []byte)
	SaveChain(data []byte)
	DeleteAllBlocks()
}

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

var b *blockchain
var once sync.Once
var dbStorage storage = db.DB{}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock() *Block {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
	return block
}

func persistBlockchain(b *blockchain) {
	dbStorage.SaveChain(utils.ToBytes(b))
}

func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Transactions...)
	}
	return txs
}

func FindTx(b *blockchain, txId string) *Tx {
	for _, tx := range Txs(b) {
		if tx.Id == txId {
			return tx
		}
	}
	return nil
}

func recalcultateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
	newestBlock := allBlocks[0]
	lastRecalcultatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalcultatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime < (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return recalcultateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)
	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Signature == "COINBASE" {
					break
				}
				if FindTx(b, input.TxId).TxOuts[input.Index].Address == address {
					creatorTxs[input.TxId] = true
				}
			}
			for index, output := range tx.TxOuts {
				if output.Address == address {
					if _, ok := creatorTxs[tx.Id]; !ok {
						uTxOut := &UTxOut{tx.Id, index, output.Amount}
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}
					}
				}
			}
		}
	}
	return uTxOuts
}

func BalanceByAddress(address string, b *blockchain) int {
	txOuts := UTxOutsByAddress(address, b)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{Height: 0}
		checkpoint := dbStorage.LoadChain()
		if checkpoint == nil {
			b.AddBlock()
		} else {
			b.restore(checkpoint)
		}
	})
	return b
}

func Status(b *blockchain, w http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()
	utils.HandleErr(json.NewEncoder(w).Encode(b))
}

func (b *blockchain) Replace(newBlocks []*Block) {
	b.m.Lock()
	defer b.m.Unlock()
	b.CurrentDifficulty = newBlocks[0].Difficulty
	b.Height = len(newBlocks)
	b.NewestHash = newBlocks[0].Hash
	persistBlockchain(b)
	dbStorage.DeleteAllBlocks()
	for _, block := range newBlocks {
		persistBlock(block)
	}
}

func (b *blockchain) AddPeerBlock(newBlock *Block) {
	b.m.Lock()
	m.m.Lock()
	defer b.m.Unlock()
	defer m.m.Unlock()
	b.Height += 1
	b.CurrentDifficulty = newBlock.Difficulty
	b.NewestHash = newBlock.Hash
	persistBlockchain(b)
	persistBlock(newBlock)

	for _, tx := range newBlock.Transactions {
		_, ok := m.Txs[tx.Id]
		if ok {
			delete(m.Txs, tx.Id)
		}
	}
}
