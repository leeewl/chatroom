package controller

import (
	"chatroom/module/chat"
	"chatroom/module/room"
	"encoding/json"
	"errors"
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
	Content     string   `json:"content"`
	Type        string   `json:"type"`
	Ip          string   `json:"ip"`
	From        string   `json:"from"`
	User        string   `json:"user"`
	Uid         string   `json:"uid"`
	Room        string   `json:"room"`      //当前房间
	RoomName    string   `json:"room_name"` //当前房间名
	UserList    []string `json:"user_list"`
	MessageList []string `json:"message_list"`
}

type connection struct {
	conn *websocket.Conn
	send chan []byte
	data *Data
	room int
	//room_list []int
}

// 管理链接的hub
type connHub struct {
	// connection 和 username
	connections map[*connection]string
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
			log.Println(err.Error())
			conn.unRegister()
			break
		}

		json.Unmarshal(message, &conn.data)

		switch conn.data.Type {
		//case typeSystem:

		case typeUser:
			conn.data.From = conn.data.User

			myroom, _ := strconv.Atoi(conn.data.Room)
			myuid, _ := strconv.Atoi(conn.data.Uid)

			chat.SaveMessage(myuid, conn.data.User, myroom, conn.data.Content)
			broadcastData, _ := json.Marshal(conn.data)
			conn.broadcast(broadcastData)
		case typeHandshake:
			// 用户连接分组时
			conn.room, err = strconv.Atoi(conn.data.Room)
			if err != nil {
				log.Printf("handshake room %v error", conn.room)
			}
			// 数据库中有分组
			rname := room.GetRoomName(conn.room)
			if rname == "" {
				// 没有分组，进入失败
				conn.handShakeFail()
			} else {
				conn.data.RoomName = rname
				conn.register()
			}
		case typeLogin:
			conn.data.Content = conn.data.User
			broadcastData, _ := json.Marshal(conn.data)
			conn.broadcast(broadcastData)
		case typeLogout:
			conn.data.Content = conn.data.User
			broadcastData, _ := json.Marshal(conn.data)
			conn.broadcast(broadcastData)
			conn.unRegister()
		case typeCreateRoom:
			// 创建分组
			// 暂时不做
		default:
			log.Println(" other type ", conn.data.Type)
		}
	}
}

func (conn *connection) write() {
	for message := range conn.send {
		err := conn.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("write error data %v conn %v ", conn.data, conn.conn)
			conn.unRegister()
			break
		}
	}
	conn.conn.Close()
}

func (conn *connection) register() error {
	if conn.room <= 0 {
		return errors.New("room id <= 0")
	}
	// 内存里没有房间
	if _, ok := rHub.roomMap[conn.room]; !ok {
		// 创建connhub
		cHub := newConnHub()
		// connHub加入roomHub
		err := rHub.addConnHub(conn.room, cHub)

		if err != nil {
			log.Println(err.Error())
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
		// string 是 username
		connections:    make(map[*connection]string),
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

func (ch *connHub) getUserList() []string {
	user_list := make([]string, 2)
	for _, v := range ch.connections {
		user_list = append(user_list, v)
	}
	return user_list
}

func (ch *connHub) run() {

	for {
		select {
		case conn := <-ch.registerConn:
			uname := conn.data.User
			ch.connections[conn] = uname
			conn.data.Type = typeHandshake
			// 从t_chat表里面找最新50条数据
			roomId, _ := strconv.Atoi(conn.data.Room)
			conn.data.MessageList = chat.GetLogInMessages(roomId)
			conn.data.UserList = ch.getUserList()
			sigleData, _ := json.Marshal(conn.data)
			conn.send <- sigleData
		case conn := <-ch.unRegisterConn:
			if _, ok := ch.connections[conn]; ok {
				delete(ch.connections, conn)
				//close(conn.send)
			}
		case data := <-ch.broadcast:
			for c := range ch.connections {
				c.send <- data
			}
		}
	}
}

// websocket的handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 日志带时间和文件
	log.SetFlags(log.Lshortfile | log.LstdFlags)

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
		close(conn.send)
	}()

	go conn.write()
	conn.read()
}

func registerWsRoute() {
	http.HandleFunc("/ws", wsHandler)
}
