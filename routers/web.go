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

	// 用户列表
	userRouter := r.Group("/user")
	{
		userRouter.GET("/list", user.List)
		userRouter.GET("/groupslist", user.GroupsList)
	}
	// home
	homeRouter := r.Group("/home")
	{
		homeRouter.GET("/index", home.Index)
		homeRouter.GET("/room", home.Room)
	}

	//ws
	r.GET("/ws",ws.WsPage)
}