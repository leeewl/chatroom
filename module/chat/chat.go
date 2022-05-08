package chat

import "time"

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
