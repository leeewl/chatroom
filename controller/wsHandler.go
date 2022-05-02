package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*
type data struct {
	Content  string   `json:"content"`
	Type     string   `json:"type"`
	Ip       string   `json:"ip"`
	From     string   `json:"from"`
	User     string   `json:"user"`
	UserList []string `json:"user_list"`
}
*/

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn.ReadMessage()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func registerWsRoute() {
	http.HandleFunc("/ws", wsHandler)
}
