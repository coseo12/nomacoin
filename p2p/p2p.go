package p2p

import (
	"fmt"
	"net/http"

	"github.com/coseo12/nomacoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("%s\n", p)
	}
}
