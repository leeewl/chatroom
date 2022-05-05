package room

type room struct {
	rid         int
	name        string
	create_time string
	create_uid  int
}

func newRoom(rid int) (r room, err error) {
	r, err = getOneRoomById(rid)
	if err != nil {
		return
	}
	return
}
