package controller

import "chatroom/config"

func RegisterRoutes(gConfig *config.GeneralConfig) {
	// css img

	// route
	registerChatRoute(gConfig)
	registerWsRoute()
	registerLoginRoute(gConfig)
	registerSignUpRoute(gConfig)
}
