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
//
//func wsPage(w http.ResponseWriter, req *http.Request) {
//	// 升级协议
//	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
//		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
//
//		return true
//	}}).Upgrade(w, req, nil)
//	if err != nil {
//		http.NotFound(w, req)
//		return
//	}
//	conn.CloseHandler()
//	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
//
//	currentTime := uint64(time.Now().Unix())
//	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)
//
//	go client.read()
//	go client.write()
//
//	// 用户连接事件
//	clientManager.Register <- client
//}
