package main

func main() {
InitConfig()




}

func InitConfig() {

	c := conf.Config{}
	c.Routes=[]string{"/ping","/renewal","/login","/login/mobile","/sendsms","/signup/mobile","/signup/mobile/exist"}
	c.OpenJwt=true//开启jwt
	conf.Set(c)
	//初始化数据验证
	handle.InitValidate()
}