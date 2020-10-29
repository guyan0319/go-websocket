package routers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/controllers/home"
	"go-websocket/controllers/user"
	"go-websocket/servers/ws"
)

//注册web服务
func Init(r *gin.Engine)  {
	r.LoadHTMLGlob("view/**/*")

	// 用户组
	userRouter := r.Group("/user")
	{
		userRouter.GET("/list", user.List)
	}

	// home
	homeRouter := r.Group("/home")
	{
		homeRouter.GET("/index", home.Index)
	}
	//ws
	r.GET("/ws",ws.WsHandler)
}