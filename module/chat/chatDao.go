package chat

import (
	"chatroom/app/injectors"
	"chatroom/infrastructure"
	"fmt"
	"log"
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

var message_num = 50

func SelectMany(room_id int) (messageSlice []Message, err error) {
	conf := injectors.GetConfig()
	db, err := infrastructure.ConDb(&conf.PostgreSqlDB)
	if err != nil {
		return
	}
	rows, err := db.Query("select uname , send_time ,message from t_chat where room_id = $1 order by send_time desc limit $2 ", room_id, message_num)
	for rows.Next() {
		m := Message{}
		err := rows.Scan(&m.uname, &m.send_time, &m.message)
		if err != nil {
			log.Fatalln(err.Error())
		}
		messageSlice = append(messageSlice, m)
	}
	return
}
