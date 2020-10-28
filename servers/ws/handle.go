package ws

import (
	"encoding/json"
	"fmt"
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

	request := &models.Request{}

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
	cmd := request.Cmd

	var (
		code uint32
		msg  string
		data interface{}
	)

	// request
	fmt.Println("acc_request", cmd, client.Addr)

	// 采用 map 注册的方式
	if value, ok := gethandlerMap(cmd); ok {
		code, msg, data = value(client, seq, requestData)
	} else {
		code = common.RoutingNotExist
		fmt.Println("处理数据 路由不存在", client.Addr, "cmd", cmd)
	}

	msg = common.GetErrorMessage(code, msg)

	responseHead := models.NewResponseHead(seq, cmd, code, msg, data)

	headByte, err := json.Marshal(responseHead)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)

		return
	}

	client.SendMsg(headByte)

	fmt.Println("acc_response send", client.Addr, client.AppId, client.UserId, "cmd", cmd, "code", code)

	return
}
