package msgs

import "go-websocket/lib/response"
const (
	MessageTypeText = "text"

	MessageActionMsg = "msg"
	MessageActionEnter = "enter"
	MessageActionExit = "exit"
)

// 消息的定义
type Message struct {
	Target string `json:"target"` // 目标
	Type   string `json:"type"`   // 消息类型 text/img/
	Msg    string `json:"msg"`    // 消息内容
	From   string `json:"from"`   // 发送者
}

func NewTestMsg(from string, Msg string) (message *Message) {

	message = &Message{
		Type: MessageTypeText,
		From: from,
		Msg:  Msg,
	}

	return
}

func getTextMsgData(action, uuId, msgId, message string) string {
	textMsg := NewTestMsg(uuId, message)
	head := NewResponseHead(msgId, action, response.OK, "Ok", textMsg)

	return head.String()
}

// 文本消息
func GetMsgData(uuId, msgId, action, message string) string {

	return getTextMsgData(action, uuId, msgId, message)
}

// 文本消息
func GetTextMsgData(uuId, msgId, message string) string {

	return getTextMsgData("msg", uuId, msgId, message)
}

// 用户进入消息
func GetTextMsgDataEnter(uuId, msgId, message string) string {

	return getTextMsgData("enter", uuId, msgId, message)
}

// 用户退出消息
func GetTextMsgDataExit(uuId, msgId, message string) string {

	return getTextMsgData("exit", uuId, msgId, message)
}
