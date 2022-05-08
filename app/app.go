package app

import (
	"chatroom/app/controller"
	"chatroom/app/injectors"
	"chatroom/config"
	"net/http"
)

type App interface {
	// 初始化
	Init() error
	// 启动 http server
	StartHttpServer() error
}

type restApiApplication struct {
	config     *config.GeneralConfig
	httpServer *http.Server
}

// 启动服务
func (app *restApiApplication) Init() error {
	// 注册路由
	controller.RegisterRoutes(app.config)
	return nil
}

// 启动http服务
func (app *restApiApplication) StartHttpServer() error {
	return app.httpServer.ListenAndServe()
}

func NewApp(configFilePaths ...string) (*restApiApplication, error) {
	generalConfig, err := injectors.ProvideConfig(configFilePaths...)
	if err != nil {
		return nil, err
	}
	server := injectors.ProvideHttpServer(generalConfig, nil)
	/*
		sqlDb, err := injectors.ProvideSqlDBConnector(generalConfig)
		if err != nil {
			return nil, err
		}
	*/
	appRestApiApplication := &restApiApplication{
		config:     generalConfig,
		httpServer: server,
	}
	return appRestApiApplication, nil
}
