package ws

import (
	"go-websocket/conf"
	"go-websocket/servers/msgs"
)

func GetAppIds() []uint32 {
	return conf.Cfg.AppIds
}

func GetServer() (server *msgs.Server) {
	server = msgs.NewServer(conf.Cfg.ServerIp, conf.Cfg.ServerPort)

	return
}

func IsLocal(server *msgs.Server) (isLocal bool) {
	if server.Ip == conf.Cfg.ServerIp && server.Port == conf.Cfg.ServerPort {
		isLocal = true
	}

	return
}

func InAppIds(appId uint32) (inAppId bool) {

	for _, value := range GetAppIds() {
		if value == appId {
			inAppId = true

			return
		}
	}

	return
}

