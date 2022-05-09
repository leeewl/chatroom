package chat

import (
	"log"
	"time"
)

type chat struct {
	cid       int
	uid       int
	uname     string
	room_id   int
	send_time int
	message   string
}

func SaveMessage(uid int, uname string, room_id int, msg string) error {

	c := chat{
		uid:       uid,
		uname:     uname,
		room_id:   room_id,
		send_time: int(time.Now().Unix()),
		message:   msg,
	}
	return Insert(c)
}

type Message struct {
	uname     string
	send_time int
	message   string
}

func GetLogInMessages(room_id int) []string {
	messages := make([]string, 2)
	messageSlice, err := SelectMany(room_id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 翻转，最新的数据在最后
	for j := len(messageSlice) - 1; j >= 0; j = j - 1 {
		messages = append(messages, messageSlice[j].uname+" : "+messageSlice[j].message)
	}
	return messages
}
