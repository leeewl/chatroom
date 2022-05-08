package chat

import (
	"chatroom/app/injectors"
	"chatroom/infrastructure"
	"fmt"
)

func Insert(c chat) (err error) {
	conf := injectors.GetConfig()
	db, err := infrastructure.ConDb(&conf.PostgreSqlDB)
	if err != nil {
		return
	}
	fmt.Println(c)
	_, err = db.Exec("Insert into t_chat (uid,uname,room_id,send_time,message) values ($1, $2,$3,$4,$5)",
		c.uid, c.uname, c.room_id, c.send_time, c.message)
	if err != nil {
		return
	}
	return
}

func Select(room_id, num int) {

}
