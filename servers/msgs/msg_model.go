package msgs

import "go-websocket/lib/response"

const (
	MessageTypeText = "text"
	MessageTypeImage = "image"
	MessageTypeVideo = "video"
	MessageTypeVoice = "voice"

	MessageActionMsg   = "sendmsg"
	MessageActionEnter = "enter"
	MessageActionExit  = "exit"
	MessageActionHeartbeat  = "heartbeat"
)

//msgTimestamp  发送时间
//msgType  消息类型 text img
//content  内容
//from  发送者
//to  接受者
// 消息的定义
type Message struct {
	Action       string `json:"action"`       //请求方法
	MsgId        string `json:"msgId"`
	MsgType      string `json:"msgType"`      // 消息类型 text/img/
	Content      string `json:"content"`      // 消息内容
	From         string `json:"from"`         // 发送者
	To           string `json:"to"`           // 接收者
	MsgTimestamp string `json:"msgTimestamp"` // 时间
}

func NewMessage(action, msgId, msgType, content, to, from, msgTimestamp string) (message *Message) {
	message = &Message{
		Action:       action,
		MsgId:        msgId,
		MsgType:      msgType,
		Content:      content,
		From:         from,
		To:           to,
		MsgTimestamp: msgTimestamp,
	}
	return
}

//func getTextMsgData(action, uuId, msgId, message string) string {
//	textMsg := NewTestMsg(uuId, message)
//	head := NewResponseHead(msgId, action, response.OK, "Ok", textMsg)
//
//	return head.String()
//}

// 文本消息
func GetMsgData(action, msgId, msgType, content, to, from, msgTimestamp string) string {
	data := NewMessage(action, msgId, msgType, content, to, from, msgTimestamp)
	res := NewResponse(response.OK, "success", data)
	return res.String()
}

//// 文本消息
//func GetTextMsgData(uuId, msgId, message string) string {
//
//	return getTextMsgData("msg", uuId, msgId, message)
//}
//
//// 用户进入消息
//func GetTextMsgDataEnter(uuId, msgId, message string) string {
//
//	return getTextMsgData("enter", uuId, msgId, message)
//}
//
//// 用户退出消息
//func GetTextMsgDataExit(uuId, msgId, message string) string {
//
//	return getTextMsgData("exit", uuId, msgId, message)
//}
