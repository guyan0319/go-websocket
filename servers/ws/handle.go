package ws

import (
	"encoding/json"
	"fmt"
	"go-websocket/lib/response"
	"go-websocket/servers/msgs"
	"sync"
)

type DisposeFunc func(client *Client, seq string, message []byte) (code uint32, msg string, data interface{})

var (
	handlerMap        = make(map[string]DisposeFunc)
	handlerMapRWMutex sync.RWMutex
)

// 注册
func Register(key string, value DisposeFunc) {
	handlerMapRWMutex.Lock()
	defer handlerMapRWMutex.Unlock()
	handlerMap[key] = value
	return
}

func gethandlerMap(key string) (value DisposeFunc, ok bool) {
	handlerMapRWMutex.RLock()
	defer handlerMapRWMutex.RUnlock()

	value, ok = handlerMap[key]

	return
}

// 处理数据
func Handle(client *Client, message []byte) {

	fmt.Println("处理数据", client.Addr, string(message))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()
	request := &msgs.Request{}

	err := json.Unmarshal(message, request)
	if err != nil {
		fmt.Println("处理数据 json Unmarshal", err)
		client.SendMsg([]byte("数据不合法"))
		return
	}

	requestData, err := json.Marshal(request.Data)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))

		return
	}

	seq := request.Seq
	action := request.Action

	var (
		code uint32
		msg  string
		data interface{}
	)

	// request
	fmt.Println("请求来源", action, client.Addr)

	// 获取所有处理方式
	if value, ok := gethandlerMap(action); ok {
		code, msg, data = value(client, seq, requestData)
	} else {
		code = response.RoutingNotExist
		fmt.Println("处理数据 路由不存在", client.Addr, "action", action)
	}

	msg = response.GetErrorMessage(code, msg)

	responseHead := msgs.NewResponseHead(seq, action, code, msg, data)

	headByte, err := json.Marshal(responseHead)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)

		return
	}

	client.SendMsg(headByte)
	fmt.Println("响应：", client.Addr, client.AppId, client.UserId, "action", action, "code", code)

	return
}
