package controller

import "chatroom/config"

func RegisterRoutes(gConfig *config.GeneralConfig) {
	// css img 有资源可以放这里

	// route 注册所有路由的地方
	registerChatRoute(gConfig)
	registerWsRoute()
	registerLoginRoute(gConfig)
	registerSignUpRoute(gConfig)
}
