package controller

import (
	"chatroom/config"
	"chatroom/module/user"
	"html/template"
	"net/http"
)

type signupHandler struct {
	generalConfig *config.GeneralConfig
}

func (sh *signupHandler) executeResult(w http.ResponseWriter, result string) {
	tmpInfo := make(map[string]interface{})
	tmpInfo["host"] = sh.generalConfig.App.HttpHost
	tmpInfo["port"] = sh.generalConfig.App.HttpPort
	tmpInfo["result"] = result

	t := template.New("result.html")
	t, _ = t.ParseFiles("app/templates/result.html")

	t.Execute(w, tmpInfo)
}

func (sh *signupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// 前端提交的数据不可信，判断
	if _, ok := r.Form["name"]; !ok {
		sh.executeResult(w, "用户名不能为空")
		return
	}
	if len(r.Form["name"]) == 0 {
		sh.executeResult(w, "用户名不能为空")
		return
	}

	if _, ok := r.Form["password"]; !ok {
		sh.executeResult(w, "密码不能为空")
		return
	}
	if len(r.Form["password"]) == 0 {
		sh.executeResult(w, "密码不能为空")
		return
	}

	uname := r.Form["name"][0]
	passwd := r.Form["password"][0]

	err := user.CreateUser(uname, passwd)
	if err != nil {
		sh.executeResult(w, "用户已存在")
	} else {
		sh.executeResult(w, "用户创建成功")
	}
}

func registerSignUpRoute(Config *config.GeneralConfig) {

	sh := signupHandler{
		generalConfig: Config,
	}
	http.Handle("/signup", &sh)
}
