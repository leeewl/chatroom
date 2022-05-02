package controller

import (
	"html/template"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/tmpl.html")
	t.Execute(w, "Hello World!")
}

func registerHelloRoute() {
	http.HandleFunc("/hello", helloHandle)
}
