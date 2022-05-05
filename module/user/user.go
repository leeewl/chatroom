package user

type user struct {
	uid           int
	uname         string
	passwd        string
	create_time   int
	ban_chat_time int
	ban_time      int
}

func newUser(uid int) (u user, err error) {
	u, err = getOneById(uid)
	if err != nil {
		return
	}
	return
}

func (u *user) getUid() int {
	return u.uid
}
