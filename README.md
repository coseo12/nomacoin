# Nomacoin

go mod init ${github}

## Packages

Packages description

- blockchain: Blockchain core

- explorer: Server side rendering in golang

- rest: REST API with mux

- utils: Used to blockchain utils

- cli: Command Line Interface with flag

- db: Interface of database with BoltDB

- wallet: Wallet Packages

- p2p: peer to peer network

## External Dependency

- [Gorilla](https://github.com/gorilla/mux)

- [Gorilla WebSocket](https://pkg.go.dev/github.com/gorilla/websocket)

- [BoltDB](https://github.com/boltdb/bolt)

- [BoltBrowser](https://github.com/br0xen/boltbrowser)

  ~/.zshrc

  ```
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOPATH/bin
  ```

- [boltdbweb](https://github.com/evnix/boltdbweb)
  boltdbweb --db-name=blockchain.db
