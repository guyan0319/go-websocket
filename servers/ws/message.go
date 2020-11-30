package ws

import (
	"fmt"
	"github.com/pkg/errors"
	"go-websocket/lib/cache"
	"go-websocket/servers/msgs"
	"time"
)

func GetMsgIdTime() (msgId string) {
	currentTime := time.Now().Nanosecond()
	msgId = fmt.Sprintf("%d", currentTime)
	return
}
func GetMsgTime() (ctime string) {
	currentTime := time.Now().Unix()
	ctime = fmt.Sprintf("%d", currentTime)
	return
}

// 给用户发送消息
func SendUserMessage(appId uint32, userId string, msgId, message string) (sendResults bool, err error) {

	//data := msgs.GetTextMsgData(userId, msgId, message)

	//// TODO::需要判断不在本机的情况
	//sendResults, err = SendUserMessageLocal(appId, userId, data)
	//if err != nil {
	//	fmt.Println("给用户发送消息", appId, userId, err)
	//}

	return
}

// 给本机用户发送消息
func SendUserMessageLocal(appId uint32, userId string, data string) (sendResults bool, err error) {

	client := GetUserClient(appId, userId)
	if client == nil {
		err = errors.New("用户不在线")
		return
	}

	// 发送消息
	client.SendMsg([]byte(data))
	sendResults = true

	return
}
// 给全体用户发消息
func SendUserMessageAll(msg *msgs.SendUserMsg,action string) (sendResults bool, err error) {
	sendResults = true
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	data := msgs.GetMsgData(action,msg.MsgId,msg.MsgType,msg.Message,msg.ToUid,msg.UserId,GetMsgTime())
	for _, server := range servers {
		if IsLocal(server) {
			SendMessages(msg.AppId, msg.ToUid, data)
		} else {
			RpcSendMessages(server,msg.AppId, msg.ToUid, data)
		}
	}
	return
}
