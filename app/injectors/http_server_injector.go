package injectors

import (
	"chatroom/config"
	"fmt"
	"net/http"
)

func ProvideHttpServer(config *config.GeneralConfig, hander http.Handler) *http.Server {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.App.HttpPort),
		Handler: hander,
	}
	return httpServer
}
