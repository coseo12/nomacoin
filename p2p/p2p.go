package p2p

import (
	"fmt"
	"net/http"
	"time"

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
		if err != nil {
			conn.Close()
			break
		}
		fmt.Printf("Just got: %s\n\n", p)
		time.Sleep(2 * time.Second)
		message := fmt.Sprintf("New Message:%s\n\n", p)
		utils.HandleErr(
			conn.WriteMessage(websocket.TextMessage, []byte(message)))
	}
}
