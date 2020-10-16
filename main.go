package main

import "go-websocket/conf"

func main() {
	//初始化配置
	InitConfig()

}

func InitConfig() {
	c := conf.Config{}
	c.Routes = []string{"/ping", "/renewal", "/login", "/login/mobile", "/sendsms", "/signup/mobile", "/signup/mobile/exist"}
	c.OpenJwt = true //开启jwt
	conf.Set(c)
}
