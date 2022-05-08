package controller

import (
	"chatroom/config"
	"html/template"
	"net/http"
)

type loginHandler struct {
	generalConfig *config.GeneralConfig
}

func (lh *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.New("login.html")
	t, _ = t.ParseFiles("app/templates/login.html")
	tmpInfo := make(map[string]interface{})
	tmpInfo["host"] = lh.generalConfig.App.HttpHost
	tmpInfo["port"] = lh.generalConfig.App.HttpPort

	t.Execute(w, tmpInfo)
}

func registerLoginRoute(Config *config.GeneralConfig) {

	lh := loginHandler{
		generalConfig: Config,
	}
	http.Handle("/login", &lh)
}
