package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	typeSystem    = "system"
	typeHandshake = "handshake"
	typeLogin     = "login"
	typeLogout    = "logout"
	typeUser      = "user"
)

// 前后端传送的数据结构
type Data struct {
	Content  string   `json:"content"`
	Type     string   `json:"type"`
	Ip       string   `json:"ip"`
	From     string   `json:"from"`
	User     string   `json:"user"`
	Room     string   `json:"room"` //当前房间
	UserList []string `json:"user_list"`
}

type connection struct {
	conn      *websocket.Conn
	send      chan []byte
	data      *Data
	room_list []int
	room_cur  int
}

type userName string

// 管理链接的hub
type connHub struct {
	connections map[*connection]userName
	// 广播
	broadcast      chan []byte
	registerConn   chan *connection
	unRegisterConn chan *connection
	mu             sync.Mutex
}

type roomHub struct {
	roomMap map[int]*connHub
	mu      sync.Mutex
}

func (conn *connection) read() {
	for {
		_, message, err := conn.conn.ReadMessage()
		// 数据错误，下线
		if err != nil {
			cHub.unRegisterConn <- conn
			break
		}

		json.Unmarshal(message, &conn.data)

		switch conn.data.Type {
		//case typeSystem:

		case typeUser:
			conn.data.From = conn.data.User
			broadcastData, _ := json.Marshal(conn.data)
			cHub.broadcast <- broadcastData
		case typeLogin:
			conn.data.Content = conn.data.User
			fmt.Println(conn.data)
			broadcastData, _ := json.Marshal(conn.data)
			cHub.broadcast <- broadcastData
		case typeLogout:
			conn.data.Content = conn.data.User
			broadcastData, _ := json.Marshal(conn.data)
			cHub.broadcast <- broadcastData
			cHub.unRegisterConn <- conn
		default:
			log.Fatalln(" other type ", conn.data.Type)
		}
	}
}

func (conn *connection) write() {
	for message := range conn.send {
		err := conn.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			cHub.unRegisterConn <- conn
			break
		}
	}
	conn.conn.Close()
}

var cHub = connHub{
	connections:    make(map[*connection]userName),
	broadcast:      make(chan []byte),
	registerConn:   make(chan *connection),
	unRegisterConn: make(chan *connection),
	mu:             sync.Mutex{},
}

var rHub = roomHub{
	mu:      sync.Mutex{},
	roomMap: make(map[int]*connHub),
}

func (ch *connHub) run() {

	for {
		select {
		case conn := <-ch.registerConn:
			uname := userName(conn.data.User)
			ch.connections[conn] = uname
			conn.data.Type = typeHandshake
			sigleData, _ := json.Marshal(conn.data)
			conn.send <- sigleData
		case conn := <-ch.unRegisterConn:
			if _, ok := ch.connections[conn]; ok {
				delete(ch.connections, conn)
				close(conn.send)
			}
		case data := <-ch.broadcast:
			for c := range ch.connections {
				c.send <- data
			}
		}
	}
}

func RunConnHub() {
	cHub.run()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in wsHandler")

	url := r.URL
	query := url.Query()

	// 只返回第一个值
	name := query.Get("name")
	room := query.Get("room")
	log.Println("name " + name + "room " + room)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn := &connection{
		conn:      ws,
		send:      make(chan []byte, 128),
		data:      &Data{},
		room_list: make([]int, 1),
		room_cur:  0,
	}

	cHub.registerConn <- conn

	defer func() {
		conn.data.Type = typeLogout
		conn.data.Content = conn.data.User
		broadcastData, _ := json.Marshal(conn.data)
		cHub.broadcast <- broadcastData
		cHub.unRegisterConn <- conn
	}()

	go conn.write()
	conn.read()
}

func registerWsRoute() {
	http.HandleFunc("/ws", wsHandler)
}
