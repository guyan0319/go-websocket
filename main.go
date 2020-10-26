package main

import (
	"github.com/gin-gonic/gin"
	"go-websocket/conf"
	"go-websocket/routers"
)

func main() {
	//初始化配置
	InitConfig()
	gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()
	// 初始化web路由
	routers.Init(r)








}

func InitConfig() {
	c := conf.Config{}
	c.Routes = []string{"/ping", "/renewal", "/login", "/login/mobile", "/sendsms", "/signup/mobile", "/signup/mobile/exist"}
	c.OpenJwt = true //开启jwt
	conf.Set(c)
}
