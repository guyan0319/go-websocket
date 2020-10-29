package ws

import (
	"fmt"
	"go-websocket/lib/cache"
	"go-websocket/models"
	"time"
)

func GetMsgIdTime() (msgId string) {
	currentTime := time.Now().Nanosecond()
	msgId = fmt.Sprintf("%d", currentTime)
	return
}

// 给全体用户发消息
func SendUserMessageAll(appId uint32, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		if IsLocal(server) {
			data := models.GetMsgData(userId, msgId, cmd, message)
			AllSendMessages(appId, userId, data)
		} else {
			//grpcclient.SendMsgAll(server, msgId, appId, userId, cmd, message)
		}
	}

	return
}