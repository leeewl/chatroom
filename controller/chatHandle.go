package controller

import (
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

	name := r.Form["name"][0]
	fmt.Println(r.Form)
	t := template.New("chat.html")
	t, _ = t.ParseFiles("templates/chat.html")
	t.Execute(w, name)
}

func registerChatRoute() {
	http.HandleFunc("/chat", chatHandle)
}
