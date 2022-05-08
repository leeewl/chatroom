package controller

import (
	"chatroom/config"
	"chatroom/module/user"
	"fmt"
	"html/template"
	"net/http"
)

type chatHandler struct {
	generalConfig *config.GeneralConfig
}

func (ch *chatHandler) executeResult(w http.ResponseWriter, result string) {
	tmpInfo := make(map[string]interface{})
	tmpInfo["ip"] = ch.generalConfig.App.HttpHost
	tmpInfo["port"] = ch.generalConfig.App.HttpPort
	tmpInfo["result"] = result

	t := template.New("result.html")
	t, _ = t.ParseFiles("app/templates/result.html")

	t.Execute(w, tmpInfo)
}

func (ch *chatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 前端提交的数据不可信，判断
	if _, ok := r.Form["name"]; !ok {
		ch.executeResult(w, "用户名不能为空")
		return
	}
	if len(r.Form["name"]) == 0 {
		ch.executeResult(w, "用户名不能为空")
		return
	}

	if _, ok := r.Form["password"]; !ok {
		ch.executeResult(w, "密码不能为空")
		return
	}
	if len(r.Form["password"]) == 0 {
		ch.executeResult(w, "密码不能为空")
		return
	}

	uname := r.Form["name"][0]
	passwd := r.Form["password"][0]

	uid, err := user.CheckUserLogin(uname, passwd)
	if err != nil {
		ch.executeResult(w, "用户名不存在或用户名密码不匹配")
		return
	}

	// 传给页面的数据
	tmpInfo := make(map[string]interface{})
	tmpInfo["uname"] = r.Form["name"][0]
	tmpInfo["uid"] = uid
	tmpInfo["ip"] = ch.generalConfig.App.HttpHost
	tmpInfo["port"] = ch.generalConfig.App.HttpPort

	fmt.Println(r.Form)
	t := template.New("chat.html")
	t, _ = t.ParseFiles("app/templates/chat.html")
	t.Execute(w, tmpInfo)
}

func registerChatRoute(gConfig *config.GeneralConfig) {

	ch := chatHandler{
		generalConfig: gConfig,
	}
	http.Handle("/chat", &ch)
}
