package controller

import (
	"chatroom/module/user"
	"fmt"
	"html/template"
	"net/http"
)

func chatHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 前端提交的数据不可信，判断
	fmt.Println(r.Form)
	if _, ok := r.Form["name"]; !ok {
		t := template.New("nouser.html")
		t, _ = t.ParseFiles("templates/nouser.html")
		t.Execute(w, nil)
		return
	}
	if len(r.Form["name"]) == 0 {
		t := template.New("nouser.html")
		t, _ = t.ParseFiles("templates/nouser.html")
		t.Execute(w, nil)
		return
	}

	uname := r.Form["name"][0]
	// 根据用户名，找到用户id
	uid := user.GetUidByUname(uname)
	if uid == 0 {
		return
	}
	// 传给页面的数据
	tmpInfo := make(map[string]interface{})
	tmpInfo["uname"] = r.Form["name"][0]
	tmpInfo["uid"] = uid

	fmt.Println(r.Form)
	t := template.New("chat.html")
	t, _ = t.ParseFiles("templates/chat.html")
	t.Execute(w, tmpInfo)
}

func registerChatRoute() {
	http.HandleFunc("/chat", chatHandle)
}
