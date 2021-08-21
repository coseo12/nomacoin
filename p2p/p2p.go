package p2p

import (
	"fmt"
	"net/http"

	"github.com/coseo12/nomacoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the request from :4000
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	utils.HandleErr(err)
	initPeer(conn, "address", "port")
}

func AddPeer(address, port string) {
	// from :4000 wants to connect to :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws", address, port), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
