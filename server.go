package main

import (
	"chatroom/app"
	"flag"
)

func main() {

	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.Parse()

	application, err := app.NewApp(configFilePath)
	if err != nil {
		panic(err)
	}

	// 初始化模块
	err = application.Init()
	if err != nil {
		panic(err)
	}

	// 启动服务
	application.StartHttpServer()

	// 注册路由
	//controller.RegisterRoutes()
	// 监听
	/*
		if err := http.ListenAndServe("127.0.0.1:18080", nil); err != nil {
			fmt.Println("err:", err)
		}
	*/
}
