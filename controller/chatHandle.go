package controller

import (
	"html/template"
	"net/http"
)

func chatHandle(w http.ResponseWriter, r *http.Request) {
	t := template.New("chat.html")
	t, _ = t.ParseFiles("templates/chat.html")
	t.Execute(w, nil)
}

func registerChatRoute() {
	http.HandleFunc("/chat", chatHandle)
}
