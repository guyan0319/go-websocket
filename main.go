package main

import (
	"github.com/gin-gonic/gin"
	"go-websocket/conf"
	"go-websocket/routers"
	"go-websocket/servers/task"
	"go-websocket/servers/ws"
)

func main() {
	gin.SetMode(gin.DebugMode) // 开发环境
	//gin.SetMode(gin.ReleaseMode) //线上环境
	//初始化配置
	conf.InitConfig()
	r := gin.Default()
	// 初始化web路由
	routers.Init(r)
	//ws路由
	routers.WebsocketInit()

	// 服务注册
	task.ServerInit()

	//启动websocket
	go ws.Manager.Start()
	//启动http
	r.Run(":"+conf.Cfg.Port) // listen and serve on 0.0.0.0:8282

}
