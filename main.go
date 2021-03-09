package main

import (
	"github.com/gin-gonic/gin"
	"go-websocket/conf"
	"go-websocket/routers"
	"go-websocket/servers/task"
	"go-websocket/servers/ws"
)

func main() {
	//初始化配置
	conf.InitConfig()
	gin.SetMode(gin.DebugMode) // 开发环境
	//gin.SetMode(gin.ReleaseMode) //线上环境
	//_,groups:=model.GetGroupsAll("1")
	//fmt.Println(groups)
	r := gin.Default()
	// 初始化web路由
	routers.Init(r)
	//ws路由
	routers.WebsocketInit()
	// 服务注册
	task.ServerInit()
	//启动websocket
	go ws.Manager.Start()
	//rpc服务启动
	go ws.RpcStart()
	//启动http
	r.Run(":"+conf.Cfg.ServerPort) // listen and serve on 0.0.0.0:8282
}
