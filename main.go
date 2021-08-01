package main

import (
	"github.com/coseo12/nomacoin/blockchain"
	"github.com/coseo12/nomacoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
