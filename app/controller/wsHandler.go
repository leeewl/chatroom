package controller

import (
	"chatroom/module/room"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	typeSystem        = "system"
	typeHandshake     = "handshake"
	typeHandshakeFail = "handshakefail"
	typeLogin         = "login"
	typeLogout        = "logout"
	typeUser          = "user"
	typeCreateRoom    = "createroom"
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
	conn *websocket.Conn
	send chan []byte
	data *Data
	room int
	//room_list []int
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
			conn.unRegister()
			break
		}

		json.Unmarshal(message, &conn.data)

		switch conn.data.Type {
		//case typeSystem:

		case typeUser:
			fmt.Printf("AAAAAA conn user\n")
			conn.data.From = conn.data.User
			broadcastData, _ := json.Marshal(conn.data)
			conn.broadcast(broadcastData)
		case typeHandshake:
			fmt.Printf("AAAAAA conn handshake %v\n", conn.data)
			// 用户连接分组时
			fmt.Printf("AAAAAA conn handshake room %v\n", conn.data.Room)
			conn.room, err = strconv.Atoi(conn.data.Room)
			if err != nil {
				log.Fatalf("handshake room %v error", conn.room)
			}
			// 数据库中有分组
			r := room.GetRoom(conn.room)
			if r == nil {
				// 没有分组，进入失败
				conn.handShakeFail()

			} else {
				conn.register()
			}
		case typeLogin:
			fmt.Printf("AAAAAA conn login\n")
			conn.data.Content = conn.data.User
			fmt.Println(conn.data)
			broadcastData, _ := json.Marshal(conn.data)
			conn.broadcast(broadcastData)
		case typeLogout:
			fmt.Printf("AAAAAA conn logout\n")
			conn.data.Content = conn.data.User
			broadcastData, _ := json.Marshal(conn.data)
			conn.broadcast(broadcastData)
			conn.unRegister()
		case typeCreateRoom:
			// 创建分组
			// 暂时不做
		default:
			log.Fatalln(" other type ", conn.data.Type)
		}
	}
}

func (conn *connection) write() {
	for message := range conn.send {
		err := conn.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			conn.unRegister()
			break
		}
	}
	conn.conn.Close()
}

func (conn *connection) register() error {
	fmt.Println("registerrrr ", conn.room)
	if conn.room <= 0 {
		return errors.New("room id <= 0")
	}
	fmt.Println("registerrrr ", rHub.roomMap)
	// 内存里没有房间
	if _, ok := rHub.roomMap[conn.room]; !ok {
		// 创建connhub
		cHub := newConnHub()
		// connHub加入roomHub
		err := rHub.addConnHub(conn.room, cHub)

		if err != nil {
			close(cHub.registerConn)
			close(cHub.unRegisterConn)
			return err
		}
		// 起一个协程监听
		go cHub.run()
	}
	rHub.roomMap[conn.room].registerConn <- conn
	return nil
}

func (conn *connection) unRegister() {
	if conn.room <= 0 {
		return
	}
	if _, ok := rHub.roomMap[conn.room]; !ok {
		return
	}
	rHub.roomMap[conn.room].unRegisterConn <- conn
}

func (conn *connection) broadcast(data []byte) {
	if conn.room <= 0 {
		return
	}
	if _, ok := rHub.roomMap[conn.room]; !ok {
		return
	}
	rHub.roomMap[conn.room].broadcast <- data

}

// 进入房间失败，通知前端
func (conn *connection) handShakeFail() {
	conn.data.Type = typeHandshakeFail
	sigleData, _ := json.Marshal(conn.data)
	conn.send <- sigleData
}

func newConnHub() (cHub *connHub) {
	cHub = &connHub{
		connections:    make(map[*connection]userName),
		broadcast:      make(chan []byte),
		registerConn:   make(chan *connection),
		unRegisterConn: make(chan *connection),
		mu:             sync.Mutex{},
	}
	return
}

var rHub = roomHub{
	mu:      sync.Mutex{},
	roomMap: make(map[int]*connHub),
}

func (rh *roomHub) addConnHub(roomId int, cHub *connHub) error {
	rh.mu.Lock()
	defer rh.mu.Unlock()
	// 加锁后再检查
	if _, ok := rHub.roomMap[roomId]; ok {
		return errors.New("roomHub have connHub")
	}
	rHub.roomMap[roomId] = cHub
	return nil
}

func (ch *connHub) run() {

	for {
		select {
		case conn := <-ch.registerConn:
			fmt.Printf("AAAAAA chub registerConn\n")
			uname := userName(conn.data.User)
			ch.connections[conn] = uname
			conn.data.Type = typeHandshake
			sigleData, _ := json.Marshal(conn.data)
			conn.send <- sigleData
		case conn := <-ch.unRegisterConn:
			fmt.Printf("AAAAAA chub unregisterConn\n")
			if _, ok := ch.connections[conn]; ok {
				delete(ch.connections, conn)
				close(conn.send)
			}
		case data := <-ch.broadcast:
			fmt.Printf("AAAAAA chub broadcast\n")
			for c := range ch.connections {
				c.send <- data
			}
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in wsHandler")

	url := r.URL
	query := url.Query()

	// 只返回第一个值
	name := query.Get("name")
	uid := query.Get("uid")
	log.Println("name " + name + "uid " + uid)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn := &connection{
		conn: ws,
		send: make(chan []byte, 128),
		data: &Data{},
		room: 0,
		//room_list: make([]int, 1),
	}

	// 进入聊天界面不注册了，进入房间才注册
	//cHub.registerConn <- conn

	defer func() {
		conn.data.Type = typeLogout
		conn.data.Content = conn.data.User
		broadcastData, _ := json.Marshal(conn.data)
		conn.broadcast(broadcastData)
		conn.unRegister()
	}()

	go conn.write()
	conn.read()
}

func registerWsRoute() {
	http.HandleFunc("/ws", wsHandler)
}
