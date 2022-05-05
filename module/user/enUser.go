package user

import "fmt"

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

func GetUidByUname(uname string) int {
	user, err := getOneByUname(uname)
	if err != nil {
		return 0
	}
	return user.uid
}
