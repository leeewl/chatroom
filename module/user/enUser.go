package user

import (
	"errors"
	"fmt"
	"time"
)

type userList map[int]*user

var users = make(userList)

func GetUser(uid int) *user {
	fmt.Println(users)
	if _, ok := users[uid]; ok {
		return users[uid]
	}

	u, err := newUser(uid)
	if err != nil {
		return nil
	}

	users[uid] = &u
	return users[uid]
}

func DelUser(uid int) bool {
	if _, ok := users[uid]; ok {
		delete(users, uid)
	}
	return true
}

func CreateUser(username, userpasswd string) error {
	u := user{
		uname:         username,
		passwd:        userpasswd,
		create_time:   int(time.Now().Unix()),
		ban_chat_time: 0,
		ban_time:      0,
	}

	return Insert(u)
}

func CheckUserLogin(username, userpasswd string) (uid int, err error) {
	u, err := getOneByUname(username)
	uid = 0

	if u == (user{}) {
		err = errors.New("用户不存在")
	}

	if u.uid == 0 {
		err = errors.New("用户不存在")
	}

	if u.passwd != userpasswd {
		err = errors.New("密码错误")
	}
	uid = u.uid
	return
}

func GetUidByUname(uname string) int {
	user, err := getOneByUname(uname)
	if err != nil {
		return 0
	}
	return user.uid
}
