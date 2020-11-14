package ws

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"go-websocket/lib/cache"
	"go-websocket/lib/response"
	"go-websocket/servers/msgs"
	"time"
)

// 用户登录
func LoginController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = response.OK
	currentTime := uint64(time.Now().Unix())
	request := &msgs.Login{}
	if err := json.Unmarshal(message, request); err != nil {
		code = response.ParameterIllegal
		fmt.Println("用户登录 解析数据失败", seq, err)
		return
	}
	if !request.CheckLogin() {
		code = response.UnauthorizedUserId
		fmt.Println("用户登录 非法的用户", seq, request.UserId)

		return
	}
	if !request.CheckToken() {
		code = response.Unauthorized
		fmt.Println("用户登录 未授权", seq, request.UserId)
		return
	}
	fmt.Println("webSocket_request 用户登录", seq, "Token", request.Token)

	if !InAppIds(request.AppId) {
		code = response.Unauthorized
		fmt.Println("用户登录 不支持的平台", seq, request.AppId)

		return
	}

	//处理登录
	client.Login(request.AppId, request.UserId, currentTime)

	// 存储数据
	userOnline := msgs.UserLogin(serverIp, serverPort, request.AppId, request.UserId, client.Addr, currentTime)
	err := cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = response.ServerError
		fmt.Println("用户登录 SetUserOnlineInfo", seq, err)

		return
	}

	// 用户登录
	login := &login{
		AppId:  request.AppId,
		UserId: request.UserId,
		Client: client,
	}
	Manager.Login <- login

	fmt.Println("用户登录 成功", seq, client.Addr, request.UserId)

	return
}

// 心跳接口
func HeartbeatController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = response.OK
	currentTime := uint64(time.Now().Unix())

	request := &msgs.HeartBeat{}
	if err := json.Unmarshal(message, request); err != nil {
		code = response.ParameterIllegal
		fmt.Println("心跳接口 解析数据失败", seq, err)

		return
	}

	fmt.Println("webSocket_request 心跳接口", client.AppId, client.UserId)

	if !client.IsLogin() {
		fmt.Println("心跳接口 用户未登录", client.AppId, client.UserId, seq)
		code = response.NotLoggedIn

		return
	}

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		if err == redis.Nil {
			code = response.NotLoggedIn
			fmt.Println("心跳接口 用户未登录", seq, client.AppId, client.UserId)

			return
		} else {
			code = response.ServerError
			fmt.Println("心跳接口 GetUserOnlineInfo", seq, client.AppId, client.UserId, err)

			return
		}
	}

	client.Heartbeat(currentTime)
	userOnline.Heartbeat(currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = response.ServerError
		fmt.Println("心跳接口 SetUserOnlineInfo", seq, client.AppId, client.UserId, err)

		return
	}

	return
}

// 发送消息
func SendUserMsgController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = response.OK
	currentTime := uint64(time.Now().Unix())

	request := &msgs.SendUserMsg{}
	if err := json.Unmarshal(message, request); err != nil {
		code = response.ParameterIllegal
		fmt.Println("发送消息接口 解析数据失败", seq, err)

		return
	}
	fmt.Println(currentTime)

	// 处理发送信息
	sendUsersMsg := &msgs.SendUserMsg{
		AppId:   client.AppId,
		UserId:  client.UserId,
		MsgId:   request.MsgId,
		Message: request.Message,
	}
	Manager.SendUserMsg <- sendUsersMsg
	return
}
