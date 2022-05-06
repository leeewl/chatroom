package user

import (
	"testing"
)

func TestUserDaoInsert(t *testing.T) {
	c := user{
		uname:         "miao",
		passwd:        "nopanic",
		create_time:   10000,
		ban_chat_time: 1000001,
		ban_time:      2000,
	}

	err := Insert(c)

	if err != nil {
		t.Errorf("User Insert Err %s", err.Error())
	}
}
