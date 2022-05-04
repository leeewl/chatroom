package main

import (
	"chatroom/controller"
	"fmt"
	"net/http"
)

func main() {
	// 注册路由
	controller.RegisterRoutes()
	go controller.RunConnHub()
	// 监听
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("err:", err)
	}
}
