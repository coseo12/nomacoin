package main

import (
	"github.com/coseo12/nomacoin/cli"
	"github.com/coseo12/nomacoin/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
