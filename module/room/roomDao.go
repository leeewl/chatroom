package room

import (
	posgre "chatroom/lib/db/postgre"
	"fmt"
)

func getOneRoomById(id int) (r room, err error) {
	r = room{}
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	row := db.QueryRow("SELECT rid,name, create_time, create_uid FROM t_room WHERE rid=$1", id)
	err = row.Scan(&r.rid, &r.name, &r.create_time, &r.create_uid)
	return
}

func getOneRoomByName(name string) (r room, err error) {
	r = room{}
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	row := db.QueryRow("SELECT rid, name, create_time, create_uid FROM t_room WHERE name=$1", name)
	err = row.Scan(&r.rid, &r.name, &r.create_time, &r.create_uid)
	return
}

func (r *room) Update() (err error) {
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	_, err = db.Exec("UPDATE t_room set name=$1 WHERE rid = $2", r.name, r.rid)
	if err != nil {
		return
	}
	return
}

// 不允许删除
/*
func Delete(id int) (err error) {
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	_, err = db.Exec("DELETE FROM t_room where rid = $1", id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
*/

func Insert(r room) (err error) {
	db, err := posgre.ConDb()
	if err != nil {
		return
	}
	fmt.Println(r)
	_, err = db.Exec("Insert into t_room (name,create_time, create_uid) values ($1, $2,$3)",
		r.name, r.create_time, r.create_uid)
	if err != nil {
		return
	}
	return
}