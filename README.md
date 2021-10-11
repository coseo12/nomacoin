# Nomacoin

go run -race main.go -mode=rest -port=4000

## Packages

Packages description

- [x] blockchain: Blockchain core

- [x] explorer: Server side rendering in golang

- [x] rest: REST API with mux

- [x] utils: Used to blockchain utils

- [x] cli: Command Line Interface with flag

- [x] db: Interface of database with BoltDB

- [x] wallet: Wallet Packages

- [x] p2p: peer to peer network

- [ ] Unit Testing

✨ Defulat

```
go test ./... -v
```

✨ Coverage

```
go test -v -coverprofile cover.out ./...
```

✨ Coverage html

```
go tool cover -html=cover.out
```

✨ Full command
```
go test -v -coverprofile cover.out ./... & go tool cover -html=cover.out
```

## Data race

- Occurs when accessing data to one data at the same time during multi-threading

## External Dependency

- [Godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

```
godoc -http=:6060
```

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

- [bbolt](https://github.com/etcd-io/bbolt)
