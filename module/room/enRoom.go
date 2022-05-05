package room

import "fmt"

type roomList map[int]*room

var rooms = make(roomList)

func GetRoom(rid int) *room {
	fmt.Println(rooms)
	if _, ok := rooms[rid]; ok {
		return rooms[rid]
	}

	r, err := newRoom(rid)
	if err != nil {
		return nil
	}

	rooms[rid] = &r
	return rooms[rid]
}

func DelRoom(rid int) bool {
	if _, ok := rooms[rid]; ok {
		delete(rooms, rid)
	}
	return true
}

func GetRidByname(name string) int {
	r, err := getOneRoomByName(name)
	if err != nil {
		return 0
	}
	return r.rid
}
