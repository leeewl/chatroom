package controller

import (
	"html/template"
	"net/http"
)

func loginHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, nil)
}

func registerLoginRoute() {
	http.HandleFunc("/login", loginHandle)
}
