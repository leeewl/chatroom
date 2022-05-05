package user

import (
	posgre "chatroom/lib/db/postgre"
	"fmt"
)

func getOneById(id int) (c user, err error) {
	c = user{}
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	row := db.QueryRow("SELECT uid,passwd, uname, create_time, ban_chat_time, ban_time FROM t_user WHERE uid=$1", id)
	err = row.Scan(&c.uid, &c.passwd, &c.uname, &c.create_time, &c.ban_chat_time, &c.ban_time)
	return
}

func getOneByUname(uname string) (c user, err error) {
	c = user{}
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	row := db.QueryRow("SELECT uid, uname, create_time, ban_chat_time, ban_time FROM t_user WHERE uname=$1", uname)
	err = row.Scan(&c.uid, &c.passwd, &c.uname, &c.create_time, &c.ban_chat_time, &c.ban_time)
	return
}

func (u *user) Update() (err error) {
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	_, err = db.Exec("UPDATE t_user set uname=$1,ban_time=$2 WHERE uid = $3", u.uname, u.ban_time, u.uid)
	if err != nil {
		return
	}
	return
}

/*
// 不允许删除
func Delete(id int) (err error) {
	db, err := posgre.ConDb()
	if err != nil {
		return
	_, err = db.Exec("DELETE FROM t_user where uid = $1", id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
*/

func Insert(u user) (err error) {
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	fmt.Println(u)
	_, err = db.Exec("Insert into t_user (uname,passwd, create_time, ban_chat_time, ban_time) values ($1, $2,$3,$4,$5)",
		u.uname, u.passwd, u.create_time, u.ban_chat_time, u.ban_time)
	if err != nil {
		return
	}
	return
}
