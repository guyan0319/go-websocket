package ws

import (
	"go-websocket/servers/msgs"
)

var (
	appIds        = []uint32{1, 2} // 全部的平台

	serverIp   string
	serverPort string
)

func GetAppIds() []uint32 {

	return appIds
}

func GetServer() (server *msgs.Server) {
	server = msgs.NewServer(serverIp, serverPort)

	return
}

func IsLocal(server *msgs.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}

	return
}

func InAppIds(appId uint32) (inAppId bool) {

	for _, value := range appIds {
		if value == appId {
			inAppId = true

			return
		}
	}

	return
}

