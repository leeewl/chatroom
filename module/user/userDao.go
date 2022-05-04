package user

import (
	"chatroom/conf"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func conDb() (err error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME)
	db, err = sql.Open(conf.DB_DRIVER, connStr)
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func getOneById(id int) (c user, err error) {
	c = user{}
	err = conDb()
	if err != nil {
		return
	}
	row := db.QueryRow("SELECT uid, uname, create_time, ban_chat_time, ban_time FROM t_user WHERE uid=$1", id)
	err = row.Scan(&c.uid, &c.uname, &c.create_time, &c.ban_chat_time, &c.ban_time)
	return
}

func getOneByUname(uname string) (c user, err error) {
	c = user{}
	err = conDb()
	if err != nil {
		return
	}
	row := db.QueryRow("SELECT uid, uname, create_time, ban_chat_time, ban_time FROM t_user WHERE uname=$1", uname)
	err = row.Scan(&c.uid, &c.uname, &c.create_time, &c.ban_chat_time, &c.ban_time)
	return
}

func (u *user) Update() (err error) {
	err = conDb()
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
	_, err = db.Exec("DELETE FROM t_user where uid = $1", id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
*/

func Insert(u user) (err error) {
	err = conDb()
	if err != nil {
		return
	}
	fmt.Println(u)
	_, err = db.Exec("Insert into t_user (uname, create_time, ban_chat_time, ban_time) values ($1, $2,$3,$4)",
		u.uname, u.create_time, u.ban_chat_time, u.ban_time)
	if err != nil {
		return
	}
	return
}
