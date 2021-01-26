package ws

import (
	"fmt"
	"github.com/pkg/errors"
	"go-websocket/lib/cache"
	"go-websocket/model"
	"go-websocket/servers/msgs"
	"strconv"
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
func SendUserMessage(msg *msgs.SendUserMsg,action string) (sendResults bool, err error) {
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
	if msg.GroupsId=="0"{
		//1对1发送
		return SendUserMessage(msg,action)
	}
	//群发送
	groupsUser:=model.GetUsersAll(msg.GroupsId)
	for _,v:=range groupsUser{
		msg.ToUid= strconv.FormatInt(v.ID,10)
		_,_=SendUserMessage(msg,action)
	}
	return true,nil
}
