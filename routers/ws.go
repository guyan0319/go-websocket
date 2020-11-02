package routers

import "go-websocket/servers/ws"

func WebsocketInit()  {
	ws.Register("login", ws.LoginController)
	//websocket.Register("heartbeat", websocket.HeartbeatController)
}

