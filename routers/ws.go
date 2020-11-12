package routers

import "go-websocket/servers/ws"

func WebsocketInit()  {
	ws.Register("login", ws.LoginController)
	ws.Register("heartbeat", ws.HeartbeatController)
	ws.Register("sendusermsg", ws.SendUserMsgController)
}

