
package task

import (
	"fmt"
	"go-websocket/lib/cache"
	"go-websocket/servers/ws"
	"runtime/debug"
	"time"
)

func ServerInit() {
	Timer(2*time.Second, 60*time.Second, server, "", serverDefer, "")
}

// 服务注册
func server(param interface{}) (result bool) {
	result = true

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务注册 stop", r, string(debug.Stack()))
		}
	}()

	server := ws.GetServer()
	currentTime := uint64(time.Now().Unix())
	//fmt.Println("服务注册", param, server, currentTime)
	cache.SetServerInfo(server, currentTime)

	return
}

// 服务下线
func serverDefer(param interface{}) (result bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务下线 stop", r, string(debug.Stack()))
		}
	}()

	fmt.Println("服务下线", param)

	server := ws.GetServer()
	cache.DelServerInfo(server)

	return
}
