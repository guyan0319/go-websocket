package conf

import (
	"go-websocket/lib/common"
	"sync"
)

type Config struct {
	Language string
	Token string
	Super string
	RedisPre string
	Host string
	ServerPort string
	RpcPort string
	OpenJwt bool
	Routes []string
	ServerIp string
	AppIds []uint32
}
var (
	Cfg  Config
	mutex   sync.Mutex
	declare sync.Once
)

func  Set(cfg Config) {
	mutex.Lock()
	Cfg = cfg
	Cfg.RedisPre=setDefault(cfg.RedisPre,"","go.websocket.redis")
	Cfg.Language=setDefault(cfg.Language,"","cn")
	Cfg.Token=setDefault(cfg.Token,"","token")
	Cfg.Super=setDefault(cfg.Super,"","admin")//超级账户
	Cfg.ServerPort=setDefault(cfg.ServerPort,"","8282")//web  ws 端口
	Cfg.RpcPort=setDefault(cfg.RpcPort,"","8287")//rpc 端口
	mutex.Unlock()
}
func setDefault( value,def ,defValue string) string {
	if value==def {
		return defValue
	}
	return value
}

func InitConfig() {
	c := Config{}
	c.Routes = []string{"/ping"}
	c.OpenJwt = true //开启jwt
	c.ServerIp=common.GetServerIp()//获取当前服务器ip
	c.AppIds=[]uint32{1, 2} //全部授权的平台
	Set(c)
}
