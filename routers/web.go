package routers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/controllers/user"
)

//注册web服务
func Init(r *gin.Engine)  {
	r.LoadHTMLGlob("views/**/*")

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
	r.GET("/ws",ws.WsHandler)
}